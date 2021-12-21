package main

import (
	"fmt"
	"log"
	"net-rpc/idl"
	"net/rpc"
)

var client *rpc.Client
var err error

func init() {
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:5656")
	if err != nil {
		log.Fatal("dial error: ", err)
	}
}

func syncCall() {
	args := &idl.Args{A: 5, B: 6}
	var reply int
	// client.Call 同步调用
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith err: ", err)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)
}

func asyncCall() {
	args := &idl.Args{A: 10, B: 5}
	ret := new(idl.Return)
	// client.GO 异步调用
	call := client.Go("Arith.Divide", args, &ret, nil)
	// call.Done 通道监听调用是否完成
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Arith: %d/%d=%d...%d\n", args.A, args.B, ret.Quo, ret.Rem)
}

func main() {
	syncCall()
	asyncCall()
}
