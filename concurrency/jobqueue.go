package main

import (
	"fmt"
	"time"
)

func worker(jobChan <-chan int) {
	for job := range jobChan {
		fmt.Println("current job:", job)
		time.Sleep(3 * time.Second)
		fmt.Println("finished job:", job)
	}
}

func enqueue(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}

func addByShareCommunicate(n int) []int {
	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, order int) {
			channel <- order
		}(channel, i)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(channel)

	return ints
}

// graceful shutdown

func jobqueue() {
	// make a channel with a capacity of 1.
	jobChan := make(chan int, 1024)

	// start the worker
	go worker(jobChan)

	// enqueue a job
	// fmt.Println("enqueue the job 1")
	// jobChan <- 1
	// fmt.Println("enqueue the job 2")
	// jobChan <- 2
	// fmt.Println("enqueue the job 3")
	// jobChan <- 3

	fmt.Println(enqueue(1, jobChan))
	fmt.Println(enqueue(2, jobChan))
	fmt.Println(enqueue(3, jobChan))

	fmt.Println("waiting the jobs")
	time.Sleep(10 * time.Second)
}
