package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/23metadatav2/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
	//// 追加自定义字段
	//newCtx := metadata.AppendToOutgoingContext(ctx, "token", "RandySun")
	md := metadata.Pairs("token", "app-test-randy")
	ctx = metadata.NewOutgoingContext(ctx, md)
	// 声明两个变量接收 响应头和最后一次响应trailer返回头
	var header, trailer metadata.MD
	res, err := grpcClient.Route(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("Call Route err:%v", err)
	}
	// 拿到响应数据之前可以先获取header
	fmt.Printf("响应前获取响应头header: %#v\n", header)
	log.Println("服务的返回响应data:", res)
	// 拿到响应数据后可以获取trailer 计算一个响应耗时多长时间统计
	fmt.Printf("响应返回头trailer: %#v\n", trailer)
	// 打印返回直

}
