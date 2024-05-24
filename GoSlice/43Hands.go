package main

import "fmt"

func main43() {
	slt := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range slt {
		fmt.Printf("%v - %v - Type %T \n", i, v, v)

	}
}
