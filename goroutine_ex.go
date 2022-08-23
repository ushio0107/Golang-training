package main

import (
	"fmt"
)

func gorou() {
	showMessage("start")
	go showMessage("1")
	go showMessage("2")
	go showMessage("3")
	go showMessage("4")
	showMessage("end")
	//var input string
	//fmt.Scanln(&input)
	fmt.Println("done")

	c := make(chan bool)
	<-c

}

func showMessage(msg string) {
	fmt.Println(msg)
}
