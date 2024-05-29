package main

import "fmt"

func E60() {
	fmt.Println("Hello Abuzar -1")
	defer fmt.Println("Hello Abuzar -2")
	fmt.Println("Hello Abuzar -3")
	defer fmt.Println("Hello Abuzar -4")
}
