/*
 * File: hashRing.go
 * Created Date: 2022-07-21 05:18:34
 * Author: ysj
 * Description:  负载均衡-一致性哈希算法
 */

package lb

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

type Node struct {
	server    string // 服务器
	weight    int    // 权重
	hashValue uint32 // 哈希value
}

type HashRing []Node // 所有节点

// 初始化
func NewHashRing(servers []string, virtualNum int) HashRing {
	weights := make([]int, len(servers))
	for i := 0; i < len(servers); i++ {
		weights[i] = 1
	}
	return NewHashRingWithWeights(servers, weights, virtualNum)
}

// 带权重初始化
func NewHashRingWithWeights(servers []string, weights []int, virtualNum int) HashRing {
	// 权重和
	weightsSum := 0.0
	for _, v := range weights {
		weightsSum += float64(v)
	}

	nodes := make([]Node, 0)

	for i := 0; i < len(servers); i++ {
		// 按权重计算节点数
		nodeNum := int(float64(virtualNum*len(servers)) * (float64(weights[i]) / weightsSum))
		for ii := 0; ii < nodeNum; ii++ {
			hashStr := fmt.Sprintf("%s:%v", servers[i], ii)
			sh := sha1.New()
			sh.Write([]byte(hashStr))
			hashKey := sh.Sum(nil)
			node := Node{
				server:    servers[i],
				weight:    weights[i],
				hashValue: getHashValue(hashKey[6:10]),
			}
			nodes = append(nodes, node)
			sh.Reset()
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].hashValue < nodes[j].hashValue
	})

	return nodes
}

func (hr HashRing) GetServer(key string) string {
	sh := sha1.New()
	sh.Write([]byte(key))
	hashKey := sh.Sum(nil)
	hashValue := getHashValue(hashKey[6:10])
	i := sort.Search(len(hr), func(i int) bool { return hr[i].hashValue >= hashValue })

	if i == len(hr) {
		i = 0
	}

	return hr[i].server
}

//将bs转成uint32
func getHashValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}
