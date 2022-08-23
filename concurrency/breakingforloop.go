package main

import (
	"fmt"
	"time"
)

func breakingforloop() {
	i := 0
	ch := make(chan string)
	go func() {
	LOOP:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <-ch:
				println(m)
				break LOOP
			default:
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ch <- "stop"
	time.Sleep(time.Second * 1)
}
