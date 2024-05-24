package main

import "fmt"

func mainX() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a
	fmt.Println("a", a)
	fmt.Println("b", b)

	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b)

	a = append(a, 8)

	fmt.Println("a", a)
	fmt.Println("b", b)

}
