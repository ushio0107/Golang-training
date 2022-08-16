package main

import (
	"fmt"
)

func main() {
	n := 3
	defer fmt.Println(n)
	n--
	defer fmt.Println(n)
	n--
	defer func() {
		fmt.Println(n)
	}()
	n--

	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	// var s1[]int //no assign a value, refernence type
	// var s2 = []int{} // have equal, empty slice
}
