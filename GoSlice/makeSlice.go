package main

import "fmt"

func mainMake() {
	// si := []string{"a", "b", "c", "d"}
	// fmt.Println(len(si))
	// fmt.Println(cap(si))

	xi := make([]int, 0, 6)

	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 99, 77, 55, 44, 33, 22, 33, 44, 55, 66, 77, 98, 76)
	fmt.Println(xi)

	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
