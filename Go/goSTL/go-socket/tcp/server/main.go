package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed. err:", err)
			break
		}
		revStr := string(buf[:n])
		fmt.Println("收到客户端发来的数据：", revStr)
		// conn.Write([]byte(revStr))
		conn.Write([]byte("信息已收到"))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed. err:", err)
		return
	}
	fmt.Println("tcp server listen on 127.0.0.1:20000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed. err:", err)
			continue
		}
		go process(conn)
	}
}
