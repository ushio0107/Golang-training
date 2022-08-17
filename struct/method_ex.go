package main

type myInt int

func (n myInt) absolute() myInt { // pass by value
	if n < 0 {
		return -n
	}
	return n
}

// func (n int) absolute() int will not work
// method must be on local type
