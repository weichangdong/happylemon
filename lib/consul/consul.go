package consul

import (
	"encoding/json"
	"fmt"
	"happylemon/conf"
	"happylemon/lib/util"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
)

type ServiceInfo struct {
	ServiceID string
	IP        string
	Port      int
	Load      int //负载
	Timestamp int //load updated ts
}
type ServiceList []ServiceInfo
type KVData struct {
	Load      int `json:"load"`
	Timestamp int `json:"ts"`
}

var (
	servics_map = sync.Map{}
)

//排序
func (l ServiceList) Len() int           { return len(l) }
func (l ServiceList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ServiceList) Less(i, j int) bool { return l[i].Load < l[j].Load }

//consul 注册grpc服务
func ConsulRegistGrpc(myconf conf.Conf) bool {
	grpcIp := util.GetLocalIp()
	consulAddr := myconf.Consul.Ip
	consulPort := myconf.Consul.Port
	grpcPort := myconf.Grpc.ApiPort
	grpcPortInt := util.Str2int(grpcPort[1:])
	r := DoRegistGrpc(consulAddr+":"+consulPort, grpcIp+grpcPort, myconf.Consul.Servername, grpcIp, grpcPortInt)
	if !r {
		panic("ConsulRegistGrpc error")
	}
	return r
}

//注册服务HTTP
// consul_addr：consul地址
// monitor_addr：健康检查地址 http://127.0.0.1:1234/check
// service_name：服务名称
// ip：ip地址
// port：端口号
func DoRegistService(consul_addr string, monitor_addr string, service_name string, ip string, port string) {
	my_service_id := service_name + "-" + ip + ":" + port

	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      my_service_id,
		Name:    service_name,
		Port:    util.Str2int(port),
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + monitor_addr + "/status",
			Interval: "30s",
			Timeout:  "1s",
		},
	}
	consulConf := api.DefaultConfig()
	consulConf.Address = consul_addr
	client, err := api.NewClient(consulConf)

	if err != nil {
		fmt.Println(err.Error() + "a1")
	}
	if err := client.Agent().ServiceRegister(service); err != nil {
		fmt.Println(err.Error() + "a2")
	}
	fmt.Printf("Registered service %q\n", service_name)
	go WaitToUnRegistService(client, my_service_id)
	return
}

//注册服务GRPC
// consul_addr：consul地址
// monitor_addr：健康检查地址 127.0.0.1:1234
// service_name：服务名称
// ip：ip地址
// port：端口号
func DoRegistGrpc(consulAddr string, monitorAddr string, serviceName string, ip string, port int) bool {
	serviceId := serviceName + "-" + ip
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			GRPC:     monitorAddr,
			Interval: "30s",
			Timeout:  "1s",
		},
	}
	consulConf := api.DefaultConfig()
	consulConf.Address = consulAddr
	client, err := api.NewClient(consulConf)
	if err != nil {
		fmt.Println(err.Error() + "a1")
		return false
	}
	if err := client.Agent().ServiceRegister(service); err != nil {
		fmt.Println(err.Error() + "a1")
		return false
	}
	fmt.Printf("Registered Grpc service %s\n", serviceName)
	go WaitToUnRegistService(client, serviceId)
	return true
}

//监听系统信号，如果服务中断或者kill时触发；服务启动是启用携程执行
//consul_client:consul API Client
//my_service_id:service ID
func WaitToUnRegistService(consul_client *api.Client, my_service_id string) {
	quit := make(chan os.Signal, 1)
	//os.Interrupt 表示中断
	//os.Kill 杀死退出进程
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	if consul_client == nil {
		return
	}
	if err := consul_client.Agent().ServiceDeregister(my_service_id); err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

//心跳检测check
//consul_addr 服务地址
//found_service 要查找的service name；target service;没有时为空
func DoDiscover(consul_addr string, found_service string) {
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			DiscoverServices(consul_addr, true, found_service)
		}
	}
}

//获取活跃的服务列表
// addr：consul服务地址
// healthyOnly:是否有心跳检测
// service_name：service_name,筛选条件，没有时为空
func DiscoverServices(addr string, healthyOnly bool, service_name string) (servics_map map[string]ServiceList) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	if err != nil {
		fmt.Println(err)
		return
	}

	services, _, err := client.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	servics_map = make(map[string]ServiceList)
	// fmt.Println("--do discover ---:", addr)

	for name := range services {
		servicesData, _, err := client.Health().Service(name, "", healthyOnly, &api.QueryOptions{})
		CheckErr(err)
		for _, entry := range servicesData {
			if service_name != entry.Service.Service && service_name != "" {
				continue
			}
			for _, health := range entry.Checks {
				if health.ServiceName != service_name && service_name != "" {
					continue
				}
				if health.ServiceName == "" {
					continue
				}
				//fmt.Println("  health nodeid:", health.Node, " service_name:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)

				var node ServiceInfo
				node.IP = entry.Service.Address
				node.Port = entry.Service.Port
				node.ServiceID = health.ServiceID

				//get data from kv store
				s := GetKeyValue(client, service_name, node.IP, node.Port)
				if len(s) > 0 {
					var data KVData
					err = json.Unmarshal([]byte(s), &data)
					if err == nil {
						node.Load = data.Load
						node.Timestamp = data.Timestamp
					}
				}
				//fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
				servics_map[health.ServiceName] = append(servics_map[health.ServiceName], node)
			}
		}
	}
	return
}

//更新自己的负载信息到相应的key
func DoUpdateKeyValue(consul_client *api.Client, consul_addr string, service_name string, ip string, port int) {
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			StoreKeyValue(consul_client, consul_addr, service_name, ip, port)
		}
	}
}

func StoreKeyValue(consul_client *api.Client, consul_addr string, service_name string, ip string, port int) {

	my_kv_key := service_name + "/" + ip + ":" + strconv.Itoa(port)

	var data KVData
	data.Load = rand.Intn(100) //暂时是随机数，到时候更新成负载
	data.Timestamp = int(time.Now().Unix())
	bys, _ := json.Marshal(&data)

	kv := &api.KVPair{
		Key:   my_kv_key,
		Flags: 0,
		Value: bys,
	}

	_, err := consul_client.KV().Put(kv, nil)
	CheckErr(err)
	fmt.Println(" store data key:", kv.Key, " value:", string(bys))
}

//获取负载信息
func GetKeyValue(consul_client *api.Client, service_name string, ip string, port int) string {
	key := service_name + "/" + ip + ":" + strconv.Itoa(port)

	kv, _, err := consul_client.KV().Get(key, nil)
	if kv == nil {
		return ""
	}
	CheckErr(err)

	return string(kv.Value)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
