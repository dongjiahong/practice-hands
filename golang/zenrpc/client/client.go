package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Param struct {
	A int `json:"a"`
	B int `json:"b"`
}
type Request struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Param  `json:"params"`
}

type Reply struct {
	Version string `json:"jsonrpc"`
	Result  int    `json:"result"`
	ID      int    `json:"id"`
}

func main() {
	// 1.链接
	conn, err := jsonrpc.Dial("http", "127.0.0.1:9999")
	if err != nil {
		log.Fatal("can't not connect to 127.0.0.1:9999")
		panic(err)
	}

	request := Request{
		ID:      1,
		Version: "2.0",
		Method:  "multiply",
		Params:  Param{A: 1, B: 4},
	}

	//reply := Reply{}
	var reply Reply

	// 2.调用
	if err := conn.Call("arith.sum", request, &reply); err != nil {
		log.Fatal("call MathService.Add error: ", err)
	}

	fmt.Printf("====> %v\n", reply)
}
