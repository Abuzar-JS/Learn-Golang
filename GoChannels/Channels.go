package main

import "fmt"

func Cha() {
	c := make(chan int, 4)
	c <- 42
	c <- 56
	c <- 20
	c <- 98

	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)

}
