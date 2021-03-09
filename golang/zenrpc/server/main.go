package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/semrush/zenrpc/v2"
	"zenrpc/jsonrpc"
)

func main() {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("arith", jsonrpc.ArithService{})
	rpc.Register("", jsonrpc.ArithService{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting arigthsrc on %s", *addr)
	log.Fatal(http.ListenAndServe(":9999", rpc))
}
