package main

import "fmt"

func mainCopy() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 40)
	copy(b, a)

	fmt.Println("a", a)
	fmt.Println("b", b)

	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b)

}
