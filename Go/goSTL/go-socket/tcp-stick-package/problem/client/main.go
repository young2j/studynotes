package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial to 127.0.0.1:3000 failed. err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "hello, long time no see."
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write msg to server failed. err:", err)
			continue
		}
	}
}
