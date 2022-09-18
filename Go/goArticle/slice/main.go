/*
 * File: main.go
 * Created Date: 2022-05-30 07:31:59
 * Author: ysj
 * Description:
 */

package main

import "fmt"

func main() {
	var slic []int
	oldCap := cap(slic)
	for i := 0; i < 2048; i++ {
		slic = append(slic, i)
		newCap := cap(slic)
		grow := float32(newCap) / float32(oldCap)
		if newCap != oldCap {
			fmt.Printf("len(slic):%v cap(slic):%v grow:%v %p\n", len(slic), cap(slic), grow, slic)
		}

		oldCap = newCap
	}
}
