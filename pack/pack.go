package pack

import "fmt"

func Adding(i, j int) int { // lower case and upper case naming have different meaning
	return i + j
}

func init() {
	fmt.Println("1 Inside package pack, Check Pack init function")
}
