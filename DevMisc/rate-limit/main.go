package main

import (
	"fmt"
	"rl/limiter"
	"time"
)

type WindowLimiter interface {
	Limit() bool
	CurrentCount() int
	Threshold() int
}

func mockRequest(n int, d time.Duration, l WindowLimiter) {
	for i := 0; i < n; i++ {
		time.Sleep(d * time.Millisecond)
		if l.Limit() {
			fmt.Printf("currentCount: %v > threshold:%v limit\n", l.CurrentCount(), l.Threshold())
		} else {
			fmt.Printf("currentCount: %v   threshold:%v pass\n", l.CurrentCount(), l.Threshold())
		}
	}
}

type BucketLimiter interface {
	Limit() bool
	CurrentCount() int
	Threshold() int
	Capacity() int
}

func mockRequest1(n int, d time.Duration, l BucketLimiter) {
	for i := 0; i < n; i++ {
		time.Sleep(d * time.Millisecond)
		if l.Limit() {
			fmt.Printf("currentCount: %v > capacity:%v limit\n", l.CurrentCount(), l.Capacity())
		} else {
			fmt.Printf("currentCount: %v   capacity:%v pass\n", l.CurrentCount(), l.Capacity())
		}
	}
}
func mockRequest2(n int, d time.Duration, l BucketLimiter) {
	for i := 0; i < n; i++ {
		time.Sleep(d * time.Millisecond)
		if l.Limit() {
			fmt.Printf("currentTokenCount: %v <= 0 limit\n", l.CurrentCount())
		} else {
			fmt.Printf("currentTokenCount: %v      pass\n", l.CurrentCount())
		}
	}
}

func main() {
	// fmt.Println("=================固定窗口算法==============")
	// fwl := limiter.NewFixedWindowLimiter(4)
	// mockRequest(10, 50, &fwl)
	// fmt.Println("------------------------------------------")
	// mockRequest(10, 500, &fwl)
	
	// fmt.Println("=================滑动窗口算法==============")
	// swl := limiter.NewSlideWindowLimiter(4, 4)
	// mockRequest(10, 40, &swl)
	// fmt.Println("------------------------------------------")
	// mockRequest(10, 70, &swl)
	
	// fmt.Println("=================漏桶算法=================")
	// lbl1 := limiter.NewLeakyBucketLimiter(4, 5)
	// mockRequest1(10, 50, &lbl1)
	// lbl2 := limiter.NewLeakyBucketLimiter(4, 5)
	// fmt.Println("------------------------------------------")
	// mockRequest1(10, 500, &lbl2)
	
	fmt.Println("=================令牌桶算法=================")
	tbl1 := limiter.NewTokenBucketLimiter(4, 5)
	mockRequest2(10, 50, &tbl1)
	fmt.Println("------------------------------------------")
	tbl2 := limiter.NewTokenBucketLimiter(4, 5)
	mockRequest2(10, 500, &tbl2)
}



