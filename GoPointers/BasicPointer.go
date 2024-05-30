package main

import (
	"fmt"
)

func Point() {
	// variable x stores value of 42
	x := 42
	y := "Hello"
	z := &y

	// Print Function Diplaying Value of X
	fmt.Println(x)

	// Print function Dislpaying location of memory where x is stored
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(&z)
	fmt.Println(*z)
	y = "Abx"

	*z = y
	fmt.Println(*z)

	// Print Statemnet shows the value of pointer x,y and type of x,y
	// fmt.Printf("%v\t%t\n", &x, &x)
	// fmt.Printf("%v\t%t\n", &y, &y)

}
