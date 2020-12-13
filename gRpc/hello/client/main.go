package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "hello/hello"

	"fmt"
)

const (
	// Address gRPC地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 链接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println("message: ", res.Message)

	grpclog.Println(res.Message)
}
