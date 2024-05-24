package main

import "fmt"

func mainA() {
	// Array literal
	a := [...]int{42, 55, 67}
	fmt.Println(a)

	// 3 dots auto decide the length of array
	b := [...]string{"Hello", "Golang"}
	fmt.Println(b)

	var c [3]int
	c[0] = 8
	c[1] = 9
	c[2] = 4
	fmt.Println(c)
	fmt.Printf("Type of a is %T \n", a)
	fmt.Printf("Type of b is %T \n", b)
	fmt.Printf("Type of c is %T \n", c)

}
