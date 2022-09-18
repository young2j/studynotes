/*
 * File: simpleRandom.go
 * Created Date: 2022-07-21 01:58:56
 * Author: ysj
 * Description:  负载均衡-简单随机算法
 */
package lb

import (
	"math/rand"
)

type SimpleRandom struct {
	servers []string
}

func NewSimpleRandom(servers []string) SimpleRandom {
	return SimpleRandom{
		servers: servers,
	}
}

func (s SimpleRandom) GetServer() string {
	n := rand.Intn(len(s.servers))
	return s.servers[n]
}
