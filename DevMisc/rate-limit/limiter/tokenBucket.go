/*
 * File: tokenBucket.go
 * Created Date: 2022-07-25 02:37:43
 * Author: ysj
 * Description:  限流算法——令牌桶算法
 */

package limiter

import (
	"fmt"
	"time"
)

type TokenBucketLimiter struct {
	capacity     int       // 桶容量
	currentCount int       // 当前令牌数
	threshold    int       // 令牌产生的速度，即每秒生成多少个令牌
	lastTime     time.Time // 上次请求的时间
}

// 初始化
func NewTokenBucketLimiter(threshold, capacity int) TokenBucketLimiter {
	return TokenBucketLimiter{
		capacity:     capacity,
		currentCount: 0,
		threshold:    threshold,
		lastTime:     time.Now(),
	}
}

// Limit
func (l *TokenBucketLimiter) Limit() bool {
	genTokenCount := int(time.Since(l.lastTime).Seconds() * float64(l.threshold))
	fmt.Printf("genTokenCount: %v ", genTokenCount)
	if genTokenCount > 0 {
		l.currentCount += genTokenCount
		l.lastTime = time.Now()
	}
	
	if l.currentCount > l.capacity {
		l.currentCount = l.capacity
	}

	if l.currentCount > 0 {
		l.currentCount--
		return false
	}

	return true
}

func (l *TokenBucketLimiter) CurrentCount() int {
	return l.currentCount
}
func (l *TokenBucketLimiter) Threshold() int {
	return l.threshold
}
func (l *TokenBucketLimiter) Capacity() int {
	return l.capacity
}
