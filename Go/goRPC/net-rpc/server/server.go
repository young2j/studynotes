package main

import (
	"fmt"
	"log"
	"net"
	"net-rpc/idl"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(idl.Arith)
	rpc.Register(arith)

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":5656")
	if err != nil {
		log.Fatal("listen err: ", err)
	}
	fmt.Println("server listen on 5656")
	http.Serve(listener, nil)
}
