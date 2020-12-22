package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup
	// lock sync.Mutex
	// 读写锁非常适合读多写少的场景
	rwLock sync.RWMutex
)

func add() {
	defer wg.Done() // 计数器-1
	for i := 0; i < 5000; i++ {

		// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁
		// lock.Lock()

		/* 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
		当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
		*/
		rwLock.Lock()
		{
			x = x + 1
		}
		// lock.Unlock()
		rwLock.Unlock()
	}
}

func main() {
	wg.Add(2) // 计数器+2
	go add()
	go add()
	wg.Wait() // 阻塞直到计数器为0
	fmt.Println(x)
}
