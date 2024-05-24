package main

import (
	"fmt"
	"math/rand"
)

func main25() {

	i := 0
	for i < 43 {
		fmt.Println("Value of i= ", i)
		x := rand.Intn(6)
		fmt.Println(x)
		switch x {
		case 0:
			fmt.Println("X is 0")
		case 1:
			fmt.Println("X is 1")
		case 2:
			fmt.Println("X is 2")
		case 3:
			fmt.Println("X is 3")
		case 4:
			fmt.Println("X is 4")
		case 5:
			fmt.Println("X is 5")
		}
		i++
	}

}
