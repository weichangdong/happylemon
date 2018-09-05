package apigrpc

import (
	"encoding/json"
	"fmt"
	"net"

	"happylemon/conf"
	"happylemon/entity"
	pb "happylemon/mygrpc/apiprotobuf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
)

type serverGrpc struct{}

func (s *serverGrpc) GrpcAct(ctx context.Context, req *pb.GrpcRequest) (*pb.GrpcResponse, error) {
	var retMsg string = entity.MsgError
	var retImg string
	var retData string
	switch req.Act {
	case "wcd":
		fmt.Println("grpc ok")
		retMsg = entity.MsgOk
	default:
		retMsg = entity.MsgError
	}
	var ret = map[string]interface{}{}
	if retMsg != entity.MsgOk {
		ret["ret"] = entity.RetCodeError
		ret["msg"] = retMsg
	} else {
		ret["ret"] = entity.RetCodeOk
		ret["msg"] = entity.MsgOk
		ret["img"] = retImg
		ret["data"] = retData
	}
	retJson, _ := json.Marshal(ret)
	return &pb.GrpcResponse{Ret: string(retJson)}, nil
}

func StartGrpc() {
	port := conf.Config.Grpc.ApiPort
	localIp := conf.Config.Grpc.ApiHost
	lis, err := net.Listen("tcp", localIp+port)
	if err != nil {
		fmt.Println("api StartGrpc error" + err.Error())
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterApiConnServer(s, &serverGrpc{})
	//新增安全检查
	healthserver := health.NewServer()
	hv1.RegisterHealthServer(s, healthserver)
	healthy := "SERVING"
	unhealthy := "NOT_SERVING"
	healthserver.SetServingStatus(healthy, hv1.HealthCheckResponse_SERVING)
	healthserver.SetServingStatus(unhealthy, hv1.HealthCheckResponse_NOT_SERVING)
	go s.Serve(lis)
}
