package main

import (
	"fmt"
	// "time"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主携程
	for i := 0; i < 2; i++ {
		// 切换任务，执行go func(),否则不会等待其执行即结束进程，除非显式sleep进行等待go func的 routine启动
		runtime.Gosched()
		fmt.Println("hello")
		// time.Sleep(1)
	}
}
