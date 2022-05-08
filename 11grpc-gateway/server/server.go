package main

import (
	"context"
	"go-grpc-example/11grpc-gateway/pkg/gateway"
	pb "go-grpc-example/11grpc-gateway/proto"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-05-08-23:40
*/

// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.InnerMessage) (*pb.OuterMessage, error) {
	res := pb.OuterMessage{
		ImportantString: "hello grpc gateway",
		Inner:           req,
	}
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(), // 校验器

		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(), // 校验器

		)),
	)
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")
	// grpc-->http
	httpServer := gateway.ProvideHTTP(Address, grpcServer)

	// 无需证书认证
	if err = httpServer.Serve(listener); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
