/*
 * File: ratelimt.go
 * Created Date: 2021-12-25 03:48:33
 * Author: ysj
 * Description:  限流中间价
 */

package ratelimit

type Limiter struct{}

func (l Limiter) Limit() bool {
	// 实现限速逻辑
	return false
}
