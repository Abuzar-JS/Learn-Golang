package main

import (
	"fmt"
)

func ForTest() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Value of i in this iteration is = %v \n", i)
		if i == 4 {
			fmt.Println("We found the value of i as 4")
		}

	}

	v := 3
	for v <= 10 {
		fmt.Printf("Value of v in this iteration is %v \n", v)
		v += 3
	}

	for c := 2; c >= 100; {
		fmt.Println("Value of c is \n", c)
		c += 10
	}

}
