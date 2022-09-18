/*
 * File: smoothPoll.go
 * Created Date: 2022-07-21 04:27:50
 * Author: ysj
 * Description:  负载均衡-平滑加权轮询算法
 */

package lb

type SmoothPoll struct {
	servers       []string
	weights       []int // 固定权重
	smoothWeights []int // 平滑权重(变动)
	weightsSum    int
}

// 初始化
func NewSmoothPoll(servers []string, weightsOpt ...[]int) SmoothPoll {
	weights := make([]int, len(servers))
	smoothWeights := make([]int, len(servers))
	// 默认权重
	for i := 0; i < len(servers); i++ {
		weights[i] = 1
	}

	// 自定义权重
	for _, wts := range weightsOpt {
		weights = wts
	}

	// 平滑权重
	copy(smoothWeights, weights)

	// 权重和
	weightsSum := 0
	for _, v := range weights {
		weightsSum += v
	}

	return SmoothPoll{
		servers:       servers,
		weights:       weights,
		smoothWeights: smoothWeights,
		weightsSum:    weightsSum,
	}
}

func (sp *SmoothPoll) GetServer() string {
	// 1. 选---选最大权重
	maxLoc := sp.findMax()
	server := sp.servers[maxLoc]

	// 2. 减---减权重和
	tempWeight := sp.smoothWeights[maxLoc] - sp.weightsSum

	// 3. 加---加各自固定权重
	for i := 0; i < len(sp.smoothWeights); i++ {
		if i == maxLoc {
			sp.smoothWeights[i] = tempWeight + sp.weights[i]
		} else {
			sp.smoothWeights[i] = sp.smoothWeights[i] + sp.weights[i]
		}
	}

	return server
}

func (sp *SmoothPoll) findMax() int {
	max := 0
	maxLoc := 0
	for i := 0; i < len(sp.smoothWeights); i++ {
		if sp.smoothWeights[i] > max {
			max = sp.smoothWeights[i]
			maxLoc = i
		}
	}
	return maxLoc
}
