package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net-rpc/idl"
)



func main()  {
	arith := new(idl.Arith)
	rpc.Register(arith)
	
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":5656")
	if err != nil {
		log.Fatal("listen err: ", err)
	}
	http.Serve(listener, nil)
}