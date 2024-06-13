package main

import "fmt"

// E60 exported to main.go
func E60() {
	fmt.Println("Hello Abuzar -1")
	defer fmt.Println("Hello Abuzar -2")
	fmt.Println("Hello Abuzar -3")
	defer fmt.Println("Hello Abuzar -4")
}
