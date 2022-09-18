/*
 * File: fixedWindow.go
 * Created Date: 2022-07-25 02:35:50
 * Author: ysj
 * Description:  限流算法——固定窗口算法
 */

package limiter

import (
	"time"
)

type FixedWindowLimiter struct {
	windowStart  time.Time // 窗口开始时间
	threshold    int       // 阈值qps,每秒请求上限
	currentCount int       // 当前请求数
}

// 初始化
func NewFixedWindowLimiter(threshold int) FixedWindowLimiter {
	return FixedWindowLimiter{
		windowStart:  time.Now(),
		threshold:    threshold,
		currentCount: 0,
	}
}

// Limit
func (l *FixedWindowLimiter) Limit() bool {
	if time.Since(l.windowStart) > 1*time.Second {
		l.windowStart = time.Now()
		l.currentCount = 0
	}
	l.currentCount++

	return l.currentCount > l.threshold
}

func (l *FixedWindowLimiter) CurrentCount() int {
	return l.currentCount
}
func (l *FixedWindowLimiter) Threshold() int {
	return l.threshold
}
