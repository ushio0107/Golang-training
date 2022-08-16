package main

import "fmt"

const (
	Monday = iota + 1 //iota means to add one to the value of every variable
	Tuesday
	Wednesday
)

func vars() { //main
	var i int = 0
	fmt.Println(i)
	fmt.Println(Monday)
	fmt.Println(Wednesday)
}
