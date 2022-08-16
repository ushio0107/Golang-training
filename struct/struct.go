package main

import "fmt"

type IPv4 []byte

func (a IPv4) String() string { //implement String()
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

type Human struct {
	name string
	age  int
}

func sturcture() { //main
	ipv4 := map[string]IPv4{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}

	for i, v := range ipv4 {
		fmt.Printf("%v : %v\n", i, v)
	}

	temp := []interface{}{}

	a := 1
	b := "foo"
	c := Human{
		name: "Akk",
		age:  21,
	}
	d := 100.1

	temp = append(temp, a, b, c, d)

	for i, value := range temp {
		if v, ok := value.(int); ok {
			fmt.Printf("%d, value: %v\n", i, v)
		} else if v, ok := value.(string); ok {
			fmt.Printf("%d, string value: %v\n", i, v)
		} else if v, ok := value.(Human); ok {
			fmt.Printf("%d, Human value: %s\n", i, v.name)
		} else {
			fmt.Printf("%d, Unknown\n", i)
		}
	}
	for i, value := range temp {
		switch v := value.(type) {
		case int:
			fmt.Printf("%d, value: %v\n", i, v)
		case string:
			fmt.Printf("%d, string value: %v\n", i, v)
		case Human:
			fmt.Printf("%d, Human value: %s\n", i, v.name)
		default:
			fmt.Printf("%d, Unknown\n", i)
		}
	}
}
