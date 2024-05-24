package main

import (
	"fmt"
)

func mainKK() {
	x := [5]int{}

	for i := 0; i < 5; i++ {
		x[i] = i
		fmt.Printf("Value of %v in this loop\n", i)
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v \n", v, v, i)
	}

}
