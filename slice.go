package main

import "fmt"

//slice

func modify(foo []string) {
	foo[1] = "c" //pass by reference
	fmt.Println("modify foo", foo)
}

func addChar(foo []string) {
	foo = append(foo, "c")
	fmt.Println("addChar: ", foo)
}

func sliceTrace() { //main
	foo := []string{"a", "b"} // 2 elements
	fmt.Println("before foo:", foo)
	modify(foo)
	fmt.Println("after foo:", foo)

	fmt.Println("--")
	// --
	fmt.Println("before foo:", foo)
	addChar(foo) // add a & mean call by reference
	fmt.Println("after foo:", foo)

	fmt.Println("--------")
	bar := foo[:1] // the first value of foo, slice
	fmt.Printf("bar address: %p\n", bar)
	fmt.Printf("foo address: %p\n", foo)
	fmt.Println("--------")

	fmt.Println("foo: ", foo)
	fmt.Println("bar:", bar)
	s1 := append(bar, "c")
	fmt.Println("s1: ", s1)
	s2 := append(bar, "d")
	fmt.Println("foo: ", foo)
	fmt.Println("s2: ", s2)
	s3 := append(bar, "e", "f")
	fmt.Println("foo: ", foo) // if add element over the length of slice, slice will not do anything (see line 18)
	fmt.Println("s2: ", s3)
}
