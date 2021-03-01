package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Golang RPC 的实现需要 5 个步骤
// 1. 定义一个服务结构
// 2. 为这个服务结构定义几个服务方法，每个方法接受两个参数和返回 error 类型
// 3. 使用 rpc.Register() 方法注册 「服务结构」 的实例
// 4. 监听套接字
// 5. 为每一个套接字调用 jsonrpc.ServerConn(conn) 方法

// 1.定义一个服务结构，因为提供的是数学服务，所以就叫MathService
type MathService struct{}

// 定义MathService所需要的参数，一般是两个int类型
type Args struct {
	Arg1, Arg2 int
}

type Reply struct {
	Res int `json:"res"`
}

// 2.实现加法服务，加法需要两个参数
// 所有jsonrpc方法只有两个参数，第一个用于接收所有参数，
// 第二个参数用于处理返回结果，是一个指针
// 所有的jsonrpc都只有一个返回值，error用于指示是否发生错误
//func (that *MathService) Add(args Args, reply *int) error {
func (that *MathService) Add(args Args, reply *Reply) error {
	//*reply = args.Arg1 + args.Arg2
	reply.Res = args.Arg1 + args.Arg2
	log.Printf("Add arg1: %d, arg2: %d, result: %+v\n", args.Arg1, args.Arg2, *reply)
	return nil
}

func main() {
	// 3.使用rpc注册
	rpc.Register(new(MathService))
	// 4.监听套接字
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	log.Println("listen tcp:9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("err: ", err)
			continue
		}
		// 5.调用jsonrpc
		go jsonrpc.ServeConn(conn)
	}
}
