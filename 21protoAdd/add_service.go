package main

import (
	"context"
	"go-grpc-example/21protoAdd/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-11-27-13:49
*/

type server struct {
	proto.UnimplementedCalcServiceServer // 兼容proto2和proto3
}

//
// Add
//  @Description: 实现加法
//  @receiver s
//  @param ctx
//  @param in
//  @return *proto.AddResponse
//  @return error
//
func (s server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	sum := in.GetX() + in.GetY()

	return &proto.AddResponse{Result: int64(sum)}, nil
}
func main() {
	// add rpc server
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("net.listen failed, err:%#v", err)
	}
	s := grpc.NewServer()
	// 注册
	proto.RegisterCalcServiceServer(s, &server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.serve failed, err:%#v", err)
		return
	}

}
