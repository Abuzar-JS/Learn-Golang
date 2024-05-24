package main

import (
	"fmt"
	"math/rand"
)

func Hand22() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("Value of x is %v \n Value of y is %v \n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Printf("%v is greater than or equal to 4 and less than or equal to 6", x)
	} else if y == 5 {
		fmt.Printf("%v is not 5 ", y)
	} else {
		fmt.Println("No cases Match")
	}

}
