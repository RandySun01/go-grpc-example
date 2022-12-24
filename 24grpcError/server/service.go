package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/24grpcError/proto"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-05-15-15:24
*/

// SimpleService 定义我们的服务
type SimpleService struct {
	pb.UnimplementedSimpleServer
	mu    sync.Mutex     // count的并发锁
	count map[string]int //存储每个name调用route的次数
}

const (
	timestampFormat = time.StampNano
	streamingCount  = 10
)

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {

	// 保证现场安全
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[req.GetData()]++ // 记录name的请求次数
	if s.count[req.GetData()] > 1 {
		// 返回请求次数限制的错误
		st := status.New(codes.ResourceExhausted, "request limit")

		// 添加错误详细信息，需要接收返回status(可自定义，本次采用谷歌)
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{
					{
						Subject:     fmt.Sprintf("name:%s", req.GetData()),
						Description: "每个name只能调用一次route",
					},
				},
			})
		if err != nil {
			// WithDetails 执行失败,返回原来st.Err()
			return nil, st.Err()
		}

		return nil, ds.Err() //带描述信息的details的error
	}
	// 正常返回
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}

	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(NetWork, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %V", err)
	}
	log.Println(Address, "net.Listing...")
	// 创建grpc服务实例

	grpcServer := grpc.NewServer()
	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{count: make(map[string]int)})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
