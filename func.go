package main

import (
	"fmt"

	"strings"
)

func add(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
} // multpie value return

func returnFunc() func() int {
	return func() int {
		return 100
	}
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
	username string
	email    string
}

func functionTraining() { //main
	a := 0
	b := 1
	fmt.Println(add(a, b))
	fmt.Println(add(3, 4))
	fmt.Println(swap(a, b))
	returnFuncType := returnFunc()
	fmt.Printf("%T\n", returnFuncType)
	fmt.Println(returnFuncType())

	func2 := func(i, j float32) float32 {
		return i + j
	}

	fmt.Printf("%T\n", func2)
	fmt.Println(func2(3.2, 5.6))

	func() {
		fmt.Println("String Print")
	}() // Anonymous function

	go func(i, j int) {
		fmt.Println(i + j)
	}(1, 2) // the keyword go is to make the function working background

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "appleboy",
	}))
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "appleboy",
		email:    "test@gmail.com",
	}))
}
