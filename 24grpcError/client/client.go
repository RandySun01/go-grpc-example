package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/24grpcError/proto"
	"log"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-05-15-15:24
*/
const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

func main() {
	// 连接服务器

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect connect: %v", err)
	}
	defer conn.Close()
	// 建立gRpc连接
	grpcClient := pb.NewSimpleClient(conn)

	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	ctx := context.Background()
	// 追加自定义字段

	// 声明两个变量接收 响应头和最后一次响应trailer返回头
	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		// 收到携带detail的error详细信息
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("QuotaFailure: %#v\n", info)
				fmt.Printf("QuotaFailure: %#v\n", info.String())
			default:
				fmt.Printf("unexpected type: %#v\n", info)
			}
		}
		log.Fatalf("Call Route err:%v", err)
	}
	// 打印返回直
	log.Println("服务的返回响应data:", res)

}
