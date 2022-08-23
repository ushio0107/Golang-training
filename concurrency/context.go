package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerCon(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func foo(ctx context.Context, name string) {
	go bar(ctx, name)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(name, "A do sth.")
		}
	}
}

func bar(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B exit")
			return
		case <-time.After(2 * time.Second):
			fmt.Println(name, "B do sth.")
		}
	}
}

func context_main() {
	ctx, cancel := context.WithCancel(context.Background())
	// go workerCon(ctx, "node 01")
	// go workerCon(ctx, "node 02")
	// go workerCon(ctx, "node 03")

	// time.Sleep(5 * time.Second)
	// fmt.Println("stop the goroutine")
	// cancel()
	// time.Sleep(5 * time.Second)

	// ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, "FooBar")
	fmt.Println("client release connection, need to notify A, B exit")
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

	const concurrencyProcesses = 10 // limit the maximum number of concurrent reading process tasks
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)
	found := make(chan int, 100)
	queue := make(chan int, 100)

	for i := 0; i < jobCount; i++ {
		queue <- i
	}
	close(queue)

	for i := 0; i < concurrencyProcesses; i++ {
		go func(queue <-chan int, found chan<- int) {
			for val := range queue {
				defer wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("job:", val, "wait time:", waitTime, "millisecond")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				found <- val
			}
		}(queue, found)
	}
	wg.Wait()
	close(found)
	var results []int
	for p := range found {
		fmt.Println("Finished job:", p)
		results = append(results, p)
	}

	fmt.Println("result:", results)
}
