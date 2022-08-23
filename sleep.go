package main

import (
	"fmt"
	"time"
)

func Print() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func doSomethingA(s string, count int, sleep int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Println("doSomethingA:", s, "loop:", i)
	}
}

func doSomethingB(s string, count int, sleep int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Println("doSomethingB:", s, "loop:", i)
	}
}

func slp() {

	go Print()
	go doSomethingA("mdfk2020", 3, 2)
	go doSomethingB("ggininder", 2, 1)

	for i := 1000; i < 1010; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
