/*
 * File: simplePoll.go
 * Created Date: 2022-07-21 03:37:45
 * Author: ysj
 * Description:  负载均衡-简单轮询算法
 */

package lb

type SimpePoll struct {
	servers []string
	pos     int
}

// 初始化
func NewSimpePoll(servers []string) SimpePoll {
	return SimpePoll{
		servers: servers,
		pos:     0,
	}
}

func (s *SimpePoll) GetServer() string {
	server := s.servers[s.pos]
	s.pos++

	if s.pos >= len(s.servers) {
		s.pos = 0
	}

	return server
}
