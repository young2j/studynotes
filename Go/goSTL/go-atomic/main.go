package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

// 普通版
func add() {
	defer wg.Done()
	x++
}

// 加锁版
func mutexAdd() {
	defer wg.Done()
	lock.Lock()
	x++
	lock.Unlock()
}

// 原子版
func atomicAdd() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	start := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		// go add() //普通版，非并发安全
		// go mutexAdd() //加锁版，并发安全，但性能开销大
		go atomicAdd() // 原子版，并发安全，性能开销优于加锁
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println("耗时：", end.Sub(start))
}
