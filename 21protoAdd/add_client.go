package main

import (
	"context"
	"flag"
	"go-grpc-example/21protoAdd/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-11-27-13:49
*/

var (
	x = flag.Int64("x", 10, "x的值")
	y = flag.Int64("y", 10, "y的值")
)

func main() {
	// 建立rpc连接 server端
	flag.Parse() // 从命令解析x 和 y
	clientConn, err := grpc.Dial("127.0.0.1:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, err:%v", err)
	}
	defer clientConn.Close()

	//创建rpc client端
	client := proto.NewCalcServiceClient(clientConn)
	// 发起rpc调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{X: int32(*x), Y: int32(*y)})
	if err != nil {
		log.Fatalf("client.Add failed, err:%v", err)
	}
	log.Printf("resp: %v	", resp.GetResult())
	// 打印结果
}
