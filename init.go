package main

import (
	"fmt"
	// _ "./pack"
	// any package you import not used, go compiler will delete it for you or the compiler will tell you the code has error.
	// Add a underscore before the package make it pass the compiler no matter you have use the package inside this code or not.
)

func init() { // it is possible to write multiple init function, but most program only have one init function inside
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

var global = convert()

func convert() int {
	return 100
}

func init() { // the init function is implemented before main function, that means, whatever you do to the variable before main, after init, it just become the setting of the function init
	global = 0
}

func initMain() {
	fmt.Println("global is: ", global)
}
