package main

import "fmt"

func E75() {

	// a := 42

	// fmt.Printf(" %v \t%v  %T\n", a, &a, &a)
	b := 30
	fmt.Printf("B Old Value = %v \t%v  %T \n", b, &b, &b)

	y := &b
	fmt.Printf("Y Old Value = %v \t%v  %T \n", y, &y, &y)

	fmt.Printf("y Value = %v \t%v  %T \n", y, *y, &y)
	fmt.Printf("B Value = %v \t%v  %T \n", b, &b, &b)

	fmt.Println(b)
	fmt.Println(*y)
}
