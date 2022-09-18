/*
 * File: leakyBucket.go
 * Created Date: 2022-07-25 02:37:07
 * Author: ysj
 * Description:  限流算法——漏桶算法
 */

package limiter

import (
	"fmt"
	"time"
)

type LeakyBucketLimiter struct {
	capacity     int       // 桶容量
	currentCount int       // 当前请求数，即当前容量
	threshold    int       // 阈值，每秒处理数
	lastTime     time.Time // 上次请求的时间
}

// 初始化
func NewLeakyBucketLimiter(threshold, capacity int) LeakyBucketLimiter {
	return LeakyBucketLimiter{
		capacity:  capacity,
		threshold: threshold,
		lastTime:  time.Now(),
	}
}

// Limit
func (l *LeakyBucketLimiter) Limit() bool {
	leakyCount := int(time.Since(l.lastTime).Seconds() * float64(l.threshold))
	fmt.Printf("leakyCount: %v ", leakyCount)
	if leakyCount > 0 {
		l.currentCount -= leakyCount
		l.lastTime = time.Now()
	}

	if l.currentCount < 0 {
		l.currentCount = 0
	}

	if l.currentCount < l.capacity {
		l.currentCount++
		return false
	}

	return true
}

func (l *LeakyBucketLimiter) CurrentCount() int {
	return l.currentCount
}
func (l *LeakyBucketLimiter) Threshold() int {
	return l.threshold
}
func (l *LeakyBucketLimiter) Capacity() int {
	return l.capacity
}
