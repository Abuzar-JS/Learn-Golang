package main

import "fmt"

func main24() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	for i := 0; i <= 100; i++ {
		main23()
	}
}
