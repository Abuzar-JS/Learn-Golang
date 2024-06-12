package main

import (
	"fmt"
)

func Ex1() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

func Ex1P2() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
