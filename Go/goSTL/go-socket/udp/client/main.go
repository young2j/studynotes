package main

import (
	"fmt"
	"net"
)

func main() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("connect to server failed. err:", err)
		return
	}
	defer udpConn.Close()

	// 发送数据
	_, err = udpConn.Write([]byte("hello, udp server."))
	if err != nil {
		fmt.Println("send data to server failed. err:", err)
		return
	}

	// 接收server返回的数据
	data := make([]byte, 4096)
	n, addr, err := udpConn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("rev data from server failed. err:", err)
		return
	}
	fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
}
