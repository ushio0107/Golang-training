package main

import (
	"fmt"
	"sync"
	"time"
)

func do(num int, wg *sync.WaitGroup) {
	fmt.Println("Start job: ", num)
	time.Sleep(1 * time.Second)
	fmt.Println("Done work ", num)
	wg.Done()
}

func syncTest() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go do(1, &wg)
	go do(2, &wg)
	go do(3, &wg)

	wg.Wait()
	fmt.Println("All Done")
}
