package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial to tcp failed. err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		//往tcp里写数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		// 接收server返回的信息
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv message from server side failed. err:", err)
			return
		}
		fmt.Println("recv from server:", string(buf[:n]))
	}
}
