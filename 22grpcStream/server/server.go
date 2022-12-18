package main

import (
	"context"
	"fmt"
	"go-grpc-example/22grpcStream/proto"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-11-27-14:50
*/
type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Reply: "Hello " + in.Name}, nil
}

// LotsOfReplies 服务端流返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *proto.HelloRequest, stream proto.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好 ",
		"hello ",
		"こんにちは ",
		"안녕하세요 ",
	}

	for _, word := range words {
		data := &proto.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfGreetings 接收客户端流式数据
func (s *server) LotsOfGreetings(stream proto.Greeter_LotsOfGreetingsServer) error {
	reply := "你好："
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv() // 接收
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += res.GetName() + " "
	}
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

// BidiHello 双向流式打招呼
func (s *server) BidiHello(stream proto.Greeter_BidiHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			log.Fatalf("服务端与客户端连接断开")
			return nil
		}
		if err != nil {
			log.Fatalf("服务端与客户端连接断开")
			return err
		}

		reply := magic(in.GetName()) // 对收到的数据做些处理

		// 返回流式响应
		if err := stream.Send(&proto.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}
func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                     // 创建gRPC服务器
	proto.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
