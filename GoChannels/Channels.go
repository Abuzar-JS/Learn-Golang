package main

import "fmt"

func Cha() {
	firstchn := make(chan int)
	fmt.Println("First chn :", firstchn)
	fmt.Printf("%T\n", firstchn)
	fmt.Printf("%v\n", firstchn)

}
