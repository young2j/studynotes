package main

import (
	"fmt"
	"net"
)

func main() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed. err:", err)
		return
	}
	defer udpConn.Close()

	for {
		// 读取发送方发送的数据
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed. err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		// 向发送方返回数据
		_, err = udpConn.WriteToUDP([]byte("信息已接收"), addr)
		if err != nil {
			fmt.Println("write to udp failed. err:", err)
			continue
		}
	}
}
