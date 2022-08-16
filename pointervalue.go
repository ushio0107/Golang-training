package main

import (
	"fmt"
	"time"
)

type car struct {
	name string
}

func (c car) SetName01(s string) { //if we modify many variable, pass by value will copy all the var, wasting memory
	fmt.Printf("1. SetName01 Address: %p\n", &c)
	c.name = s
}

func (c *car) SetName02(s string) { //mostly if the code have struct part, the function will write like this form
	fmt.Printf("2. SetName02 Address: %p\n", c)
	c.name = s
}

type email struct {
	from string
	to   string
}

func (e email) FromWho(s string) email { //****
	e.from = s
	return e
}

// func (e *email) FromWho(s string) { //****
//		e.from = s
//	}
//

func (e email) FromTo(s string) email {
	e.to = s
	return e
}

func (e email) Send() {
	fmt.Printf("From: %s, To: %s\n", e.from, e.to)
}

func main() { //main
	c := &car{}

	fmt.Printf("C Address: %p\n", c) //*car, not need &
	c.SetName01("foo")               //pass by value
	fmt.Println(c.name)
	fmt.Printf("C Address: %p\n", c)
	c.SetName02("bar") // pass by reference
	fmt.Println(c.name)

	num := 10
	p := &num
	fmt.Println("num address ", p)
	fmt.Println("num value ", *p)

	e := &email{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			e.FromWho(fmt.Sprintf("example%02d@example.com", i)).
				FromTo(fmt.Sprintf("example%02d@example.com", i+1)).
				Send()
		}(i)
	}

	time.Sleep(1 * time.Second)
}
