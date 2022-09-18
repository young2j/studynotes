/*
 * File: main.go
 * Created Date: 2022-07-21 01:10:52
 * Author: ysj
 * Description: 负载均衡
 */

package main

import (
	"fmt"
	"lb/lb"
)

func sampling(n int, getServer func() string) {
	out := make(map[string]float32)
	for i := 0; i < n; i++ {
		server := getServer()
		out[server] += 1.
	}

	for k, v := range out {
		fmt.Printf("%v: %v %.0f%%\n", k, v, v/float32(n)*100)
	}
}

func hashSampling(n int, f func(arg string) string) {
	out := make(map[string]float32)
	for i := 0; i < n; i++ {
		server := f(fmt.Sprintf("%v", i))
		out[server] += 1.
	}

	for k, v := range out {
		fmt.Printf("%v: %v %.0f%%\n", k, v, v/float32(n)*100)
	}
}

func main() {
	servers := []string{
		"192.168.10.1",
		"192.168.10.2",
		"192.168.10.3",
		"192.168.10.4",
		"192.168.10.5",
	}
	fmt.Println("==========简单随机==========")
	sr := lb.NewSimpleRandom(servers)
	sampling(100000, sr.GetServer)

	fmt.Println("==========加权随机(权重1)==========")
	wr1 := lb.NewRandom(servers)
	sampling(100000, wr1.GetServer)

	fmt.Println("==========加权随机==========")
	wr2 := lb.NewRandom(servers, []int{1, 1, 2, 2, 4})
	sampling(100000, wr2.GetServer)

	fmt.Println("==========简单轮询==========")
	sp := lb.NewSimpePoll(servers)
	sampling(100000, sp.GetServer)

	fmt.Println("==========加权轮询(权重1)==========")
	wp1 := lb.NewPoll(servers)
	sampling(100000, wp1.GetServer)

	fmt.Println("==========加权轮询==========")
	wp2 := lb.NewPoll(servers, []int{1, 1, 6, 1, 1})
	sampling(100000, wp2.GetServer)
	for i := 0; i < 10; i++ {
		server := wp2.GetServer()
		fmt.Printf("server: %v\n", server)
	}

	fmt.Println("==========平滑加权轮询(权重1)==========")
	smp1 := lb.NewSmoothPoll(servers)
	sampling(100000, smp1.GetServer)

	fmt.Println("==========平滑加权轮询==========")
	smp2 := lb.NewSmoothPoll(servers, []int{1, 1, 6, 1, 1})
	sampling(100000, smp2.GetServer)

	for i := 0; i < 10; i++ {
		server := smp2.GetServer()
		fmt.Printf("server: %v\n", server)
	}

	fmt.Println("==========一致性hash(等权)==========")
	hr1 := lb.NewHashRing(servers, 20)
	hashSampling(100000, hr1.GetServer)

	fmt.Println("==========一致性hash==========")
	hr2 := lb.NewHashRingWithWeights(servers, []int{1, 1, 6, 1, 1}, 20)
	hashSampling(100000, hr2.GetServer)
}
