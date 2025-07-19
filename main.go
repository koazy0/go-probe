package main

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	_ "github.com/koazy0/go-probe/global"
	"github.com/koazy0/go-probe/model"
	pb "github.com/koazy0/go-probe/proto"
	"github.com/koazy0/go-probe/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	proxy.SetUp()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterControlServiceServer(grpcServer, &controlServer{})
	log.Println("gRPC 服务器已启动，监听端口 :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}

type controlServer struct {
	pb.UnimplementedControlServiceServer
}

func (s *controlServer) PushRules(ctx context.Context, req *pb.PushRulesRequest) (*pb.Response, error) {
	for _, rule := range req.Rules {
		log.Printf("接收到规则：SourceAddr=%s, DestinationAddr=%s", rule.SourceAddr, rule.DestinationAddr)
	}
	rules := []model.ProxyRule{}
	err := gconv.Struct(req.Rules, &rules)
	if err != nil {
		log.Println(err.Error())
		return &pb.Response{
			Code:    500,
			Message: "internal error",
		}, nil
	}
	proxy.ProxyChan <- rules
	return &pb.Response{
		Code:    0,
		Message: "成功",
	}, nil
}
