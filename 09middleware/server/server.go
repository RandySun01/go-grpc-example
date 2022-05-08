package main

import (
	"context"
	"go-grpc-example/09middleware/pkg/middleware/auth"
	"go-grpc-example/09middleware/pkg/middleware/cred"
	"go-grpc-example/09middleware/pkg/middleware/recovery"
	"go-grpc-example/09middleware/pkg/middleware/zap"
	pb "go-grpc-example/09middleware/proto"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-05-08-18:56
*/
// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	// 添加拦截器后，方法里省略Token认证
	// //检测Token是否有效
	// if err := Check(ctx); err != nil {
	// 	return nil, err
	// }
	l := ctxzap.Extract(ctx)
	l.Info("年后")
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	panic("RecoveryInterceptor very bad thing happened")
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
	//opts := []grpc_ctxtags.Option{
	//	grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.TagBasedRequestFieldExtractor("log_fields")),
	//}
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			// grpc_ctxtags.StreamServerInterceptor(),
			// grpc_opentracing.StreamServerInterceptor(),
			// grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			//grpc_ctxtags.UnaryServerInterceptor(opts...),
			// grpc_opentracing.UnaryServerInterceptor(),
			// grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)

	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
