package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"resolve/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed. err:", err)
			return
		}
		fmt.Println("recv msg from client:", msg)
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen tcp 127.0.0.1:30000 failed. err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed. err:", err)
			continue
		}
		go process(conn)
	}

}
