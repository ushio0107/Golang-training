package main

import (
	"fmt"

	"github.com/go-training/pack"
)

func init() {
	fmt.Println("2 Inside executionOrder.go, Check execution Order init function")
}

func main() {
	a := 10
	b := 20
	c := a + b
	d := pack.Adding(a, b)

	fmt.Println(c, d)
}
