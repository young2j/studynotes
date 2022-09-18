/*
 * File: weightRandom.go
 * Created Date: 2022-07-21 02:21:15
 * Author: ysj
 * Description:  负载均衡-简单与加权随机算法
 */

package lb

import "math/rand"

type Random struct {
	servers   []string
	weights   []int
	weightSum int
}

// 初始化
func NewRandom(servers []string, weightsOpt ...[]int) Random {
	weights := make([]int, len(servers))
	// 默认权重
	for i := 0; i < len(servers); i++ {
		weights[i] = 1
	}

	// 自定义权重
	for _, wts := range weightsOpt {
		weights = wts
	}

	// 权重和
	weightSum := 0
	for _, v := range weights {
		weightSum += v
	}

	return Random{
		servers:   servers,
		weights:   weights,
		weightSum: weightSum,
	}
}

func (r Random) GetServer() string {
	wt := rand.Intn(r.weightSum)+1

	for i, s := range r.servers {
		if wt <= r.weights[i] {
			return s
		}
		wt -= r.weights[i]
	}

	return r.servers[0]
}
