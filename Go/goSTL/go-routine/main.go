package main

import (
	"fmt"
	"math/rand"
)

/* 求随机数各位数字之和*/

// Job struct
type Job struct {
	ID      int
	RandNum int
}

// Result struct
type Result struct {
	job *Job
	sum int
}

func main() {
	// 通道
	jobChan := make(chan *Job, 128)
	resChan := make(chan *Result, 128)

	// 消费
	createPool(64, jobChan, resChan)

	// 打印协程
	go func(resChan chan *Result) {
		for result := range resChan {
			fmt.Printf("job id: %v randnum: %v result: %v\n", result.job.ID, result.job.RandNum, result.sum)
		}
	}(resChan)

	// 不断生产
	var id int
	for {
		id++
		job := &Job{
			ID:      id,
			RandNum: rand.Int(),
		}
		jobChan <- job
	}
}

// createPool num协程个数
func createPool(num int, jobChan chan *Job, resChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resChan chan *Result) {
			for job := range jobChan {
				randNum := job.RandNum
				var sum int
				for randNum != 0 {
					num := randNum % 10
					sum += num
					randNum /= 10
				}
				result := &Result{
					job: job,
					sum: sum,
				}
				resChan <- result
			}
		}(jobChan, resChan)
	}
}
