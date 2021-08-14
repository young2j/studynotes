/*
 * Created Date: 2021-08-15 02:39:42
 * Author: ysj
 * Description: '并发入坑'
 */

package main

import (
	"fmt"
	"sync"
)

var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results  []int
	lock  = sync.RWMutex{}
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			defer wg.Done()
			res := task(num)
			lock.Lock()
			results = append(results, res)
			lock.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(results)
}
