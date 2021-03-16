package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	request := Request{
		ID:      1,
		Version: "2.0",
		Method:  "multiply",
		Params:  Param{A: 2, B: 4},
	}

	jsonStr, _ := json.Marshal(&request)

	resp, err := http.Post("http://127.0.0.1:9999", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Panicln(err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("body: %v\n", string(b))

}
