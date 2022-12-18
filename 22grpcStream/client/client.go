package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"go-grpc-example/22grpcStream/proto"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-11-27-14:50
*/

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:9999", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	// 普通rpc调用
	//sayHello(c)

	// 调用服务端流式rpc
	//callLotsOfReplies(c)

	// 客户端流式rpc
	//runLotsOfGreeting(c)

	// 双向流
	runBidiHello(c)
}

// 普通rpc调用
func sayHello(c proto.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}

// 调用服务端流式rpc
func callLotsOfReplies(c proto.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Println("c.LotsOfReplies err", err)
		return
	}
	// 依次从流式响应中读取返回的响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv failed, err: %#v\n", err)
			return
		}
		log.Printf("recv: %#v\n", res.GetReply())

	}

}

// 客户端流
func runLotsOfGreeting(c proto.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式RPC
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
	}
	names := []string{"Randy", "张三", "bary"}
	for _, name := range names {
		// 发送流式数据
		err := stream.Send(&proto.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err: %v", name, err)
		}
	}
	// 流发送结束之后关闭流并获取响应结果
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed: %v", err)
	}
	log.Printf("got reply: %v", res.GetReply())
}

// 双向流
func runBidiHello(c proto.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&proto.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
