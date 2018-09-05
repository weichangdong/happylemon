package apigrpc

import (
	"fmt"
	"happylemon/conf"
	"happylemon/lib/log"
	pb "happylemon/mygrpc/apiprotobuf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func DialGrpc(act string, para string) {
	port := conf.Config.Grpc.ApiPort
	localIp := conf.Config.Grpc.ApiHost
	address := localIp + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.ErrorLog("did not api connect: " + err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewApiConnClient(conn)
	if act == "" {
		act = "wcd"
	}
	log.InfoLog("Grpc-Reuest-data: Act: " + act + "; Para: " + para)
	r, err := c.GrpcAct(context.Background(), &pb.GrpcRequest{Act: act, Para: para})
	if err != nil {
		log.ErrorLog("did not api greet: " + err.Error())
		return
	}
	log.InfoLog("Grpc-Response-data: " + r.Ret)
	fmt.Printf("response: %s", r.Ret)
}
