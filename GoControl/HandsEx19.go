package main

import (
	"fmt"
	"math/rand"
)

func HandEx19() {
	var x int = rand.Intn(250)
	fmt.Println("x = ", x)

	if x <= 100 {
		fmt.Println(" between 0 and 100")
	} else if x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x <= 250 {
		fmt.Println("between 201 and 250")
	}

	for a := 1; a <= 10; a++ {
		fmt.Println(rand.Intn(3))
	}

}
