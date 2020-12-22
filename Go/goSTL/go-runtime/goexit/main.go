package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A.defer") //A
		func() {
			defer fmt.Println("B.defer") //B
			// 结束协程====>结束后依次往回延迟执行B,A
			// 结果是 B.defer A.defer
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
