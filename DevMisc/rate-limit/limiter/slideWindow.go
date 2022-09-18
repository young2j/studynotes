/*
 * File: slideWindow.go
 * Created Date: 2022-07-25 02:36:07
 * Author: ysj
 * Description:  限流算法——滑动窗口算法
 */

package limiter

import (
	"time"
)

type SlideWindowLimiter struct {
	windowSize      int           // 窗格数
	interval        time.Duration // 每个窗格时间间隔
	windows         []int         // 每个窗格请求数
	currentWindowId int           // 当前窗格id
	lastTime        time.Time     // 上次请求的时间
	currentCount    int           // 当前请求数
	threshold       int           // 阈值qps,每秒请求上限
}

// 初始化
func NewSlideWindowLimiter(threshold, windowSize int) SlideWindowLimiter {

	return SlideWindowLimiter{
		windowSize:      windowSize,
		interval:        time.Second / time.Duration(windowSize),
		windows:         make([]int, windowSize),
		currentWindowId: 0,
		lastTime:        time.Now(),
		currentCount:    0,
		threshold:       threshold,
	}
}

// Limit
func (l *SlideWindowLimiter) Limit() bool {

	if time.Since(l.lastTime) > l.interval {
		l.currentCount -= l.windows[l.currentWindowId]
		l.currentWindowId++
		l.currentWindowId = l.currentWindowId % l.windowSize
		l.windows[l.currentWindowId] = 0
		l.lastTime = time.Now()
	}

	l.currentCount++
	l.windows[l.currentWindowId]++

	return l.currentCount > l.threshold
}

func (l *SlideWindowLimiter) CurrentCount() int {
	return l.currentCount
}
func (l *SlideWindowLimiter) Threshold() int {
	return l.threshold
}
