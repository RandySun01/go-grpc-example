package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/23metadatav2/proto"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

/*
@author RandySun
@create 2022-05-15-15:24
*/

// SimpleService 定义我们的服务
type SimpleService struct {
	pb.UnimplementedSimpleServer
}

const (
	timestampFormat = time.StampNano
	streamingCount  = 10
)

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	// 获取请求头中消息
	// 在执行业务逻辑之前 check metadata中是否包含token
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("metadate md: %#v, ok: %#v\n", md, ok)
	if !ok {
		// 没有元数据拒绝连接
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}
	// 校验token
	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-randy" {
		return nil, status.Error(codes.Unauthenticated, "无效的token")
	}
	//if value, ok := md["token"]; ok {
	//	if len(value) > 0 && value[0] == "app-test-randy" {
	//		// 有效请求
	//	}
	//
	//}

	log.Printf("md: %+v\n", md)
	// 验证成功发送数据
	// 返回头响应添加header返回响应
	header := metadata.New(map[string]string{
		"location": "shanghai",
	})
	grpc.SetHeader(ctx, header)
	// 利用 defer 在发送完响应数据知州，发送trailer 处理完成响应最后一个响应返回trailer(流会返回多次响应)
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()
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
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
