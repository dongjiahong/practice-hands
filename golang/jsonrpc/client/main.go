package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 客户端调用 jsonrpc 有两个步骤
// 1. 使用 jsonrpc.Dial() 方法连接到服务，并返回一个连接 conn
// 2. 调用 conn.Call() 方法调用服务

// 定义MathService所需要的参数，一般是两个int类型
type Args struct {
	a, b int
}

type Reply struct {
	Res int `json:"res"`
}

func main() {
	// 1.链接
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal("can't not connect to 127.0.0.1:9090")
		panic(err)
	}

	//reply := Reply{}
	var reply Reply
	//var reply int
	var args = Args{15, 3}

	// 2.调用
	if err := conn.Call("Sum", args, &reply); err != nil {
		log.Fatal("call MathService.Add error: ", err)
	}

	fmt.Printf("MathService.Add(%d, %d)=%+v\n", args.a, args.b, reply)
}
