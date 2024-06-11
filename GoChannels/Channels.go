package main

import "fmt"

func AddtoChannel(a, b int) int {
	return a + b
}

func Cha() {
	ch := make(chan int)
	fmt.Println("Hello ")

	ch <- 3
	//fmt.Println(<-firstchn)
	go fmt.Println(AddtoChannel(6, 2))

}
