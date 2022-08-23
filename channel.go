package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, index *int) {
	sum := 0
	fmt.Printf("%d, In sum function\n", *index)
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	*index = *index + 1
}

func sum2(s []int, index *int) int {
	sum := 0
	fmt.Printf("%d, In sum2 function\n", *index)
	for _, v := range s {
		sum += v
	}
	*index = *index + 1
	return sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 0) // unbuffered channel
	index := 1
	go sum(s[:len(s)/2], c, &index)
	fmt.Println("goroutine sum1")
	go sum(s[len(s)/2:], c, &index)
	//fmt.Println(index)
	fmt.Println("goroutine sum2")
	sum2(s[:len(s)/2], &index)
	go sum2(s[:len(s)/2], &index)
	//fmt.Println(index)
	//x, y := <-c, <-c // receive from c
	//fmt.Println(x, y, x+y)

	timeout := make(chan bool, 1)
	go func() {
		fmt.Println("1 In go routine function")
		time.Sleep(1 * time.Second)
		fmt.Println("2 In go routine function")
		timeout <- true
	}()
	go func() {
		fmt.Println("go routine ")
		//time.Sleep(1 * time.Second)
	}()
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 02")
	}

	// buffer channel
	bufch := make(chan int, 2)
	bufch <- 1
	select {
	case bufch <- 2:
		fmt.Println("1channel value is ", <-bufch)
		fmt.Println("2channel value is ", <-bufch)
	default:
		fmt.Println("channel blocking")
	}

	//breaking for loop
	//i := 0
	//breakch
}

// goroutine function -> will not wait for it return
