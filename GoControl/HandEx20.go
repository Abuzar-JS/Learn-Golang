package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization of my porgram occurs")
}

func HandEx20() {
	var x int = rand.Intn(250)
	fmt.Println("x = ", x)

	switch {
	case x <= 100:
		fmt.Println(" between 0 and 100")
	case x <= 200:
		fmt.Println("Between 101 and 200")
	case x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}

	// for a := 1; a <= 10; a++ {
	// 	fmt.Println(rand.Intn(3))
	// }

}
