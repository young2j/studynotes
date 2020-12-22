package main

import (
	"fmt"
	"net"
	"resolve/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial tcp 127.0.0.1:30000 failed. err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "hello server,long time no see. "
		data, err := proto.Encode(msg)
		fmt.Println("send msg:", data)
		if err != nil {
			fmt.Println("encode msg failed. err:", err)
			return
		}
		conn.Write(data)
	}
}
