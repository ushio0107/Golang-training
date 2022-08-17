package main

import (
	"fmt"
	"reflect"
)

// type in GO is like typedef in C
type myint int

type person struct {
	name   string
	age    uint // unsigned int
	groups []string
	crush  *person
}

func addOne(p *person) { //pass by reference
	p.age = p.age + 1 // = (*p).age
	fmt.Println(p.age)
}

// func
func GetName(p *person) string {
	return p.name
}

// method
func (p *person) GetName() string {
	return p.name
}

func (p *person) getGroups() []string {
	return p.groups
}

func (p *person) setCrush(crush *person) {
	p.crush = crush
}

func (p *person) getCrushName() string {
	return p.crush.name
}

// declare a variable outside func, main or anythings means that this variable

func structex() {

	// declare
	var toby person
	emma := &person{}
	molly := person{
		name: "molly",
		age:  22,
	}
	kelly := new(person)
	emma.name = "emma"
	// fmt.Println(reflect.TypeOf(toby)) // output main.person

	toby.name = "toby"
	toby.age = 22
	addOne(&toby) // +& -> get the address of toby

	// similar to malloc in C language
	kelly.name = "kelly"
	kelly.age = 10

	fmt.Printf("kelly address: %p\n", kelly)
	fmt.Println(reflect.TypeOf(kelly)) // output *main.person, new(pointer)
	fmt.Println(kelly)

	fmt.Println("--------")
	fmt.Println("molly: ", molly)
	// output &{kelly 10}, print & at the beginning means the output is token from the value of the address,
	// print out the content of the memory
	fmt.Println(toby.GetName())
	fmt.Println(GetName(&toby))
	fmt.Println(GetName(kelly)) // different declare
	// var toby person, var kelly *person

}
