package main

import "fmt"

func main46() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(x[0:4])
	// fmt.Println(x[2:7])
	x = append(x[:3], x[5:]...)
	fmt.Println(x)
}
