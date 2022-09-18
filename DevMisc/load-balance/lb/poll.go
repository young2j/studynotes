/*
 * File: poll.go
 * Created Date: 2022-07-21 03:56:22
 * Author: ysj
 * Description:  负载均衡-简单与加权轮询
 */

package lb

type Poll struct {
	servers   []string
	weights   []int
	weightSum int
	index     int64
}

// 初始化
func NewPoll(servers []string, weightsOpt ...[]int) Poll {
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
	weightsSum := 0
	for _, v := range weights {
		weightsSum += v
	}

	return Poll{
		servers:   servers,
		weights:   weights,
		weightSum: weightsSum,
	}
}

func (p *Poll) GetServer() string {
	p.index++
	pos := p.index % int64(p.weightSum)
	for i, s := range p.servers {
		if pos < int64(p.weights[i]) {
			return s
		}
		pos -= int64(p.weights[i])
	}

	return p.servers[0]
}
