package main

import "fmt"

func Ifunc() {
	a := 422
	b := 932

	if a == b {
		fmt.Println("A is equal to B")
	} else if a >= b {
		fmt.Println("A is greater than B")

	} else if a <= b {
		fmt.Println("A is less than b")

	} else {
		fmt.Println("no condition Satisfies")
	}
}
