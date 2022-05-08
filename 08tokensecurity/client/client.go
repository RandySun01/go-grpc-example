package main

import (
	"context"
	"go-grpc-example/08tokensecurity/pkg/auth"
	pb "go-grpc-example/08tokensecurity/proto"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-05-08-17:29
*/
// Address 连接地址
const Address string = ":8001"

var grpcClient pb.SimpleClient

func main() {
	//从输入的证书文件中为客户端构造TLS凭证

	//构建Token
	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "123456",
	}
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	// 每次请求方法都会调用 grpc.WithPerRPCCredentials(&token) 获取token
	route()
	route()
}

// route 调用服务端Route方法
func route() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
