package main

import (
	"fmt"
	"math/rand"
)

func mainY() {
	fmt.Println("Hello Lets Start")
	a := 32
	b := 40
	fmt.Printf(" a = %v \n b = %v \n", a, b)
	fmt.Println(rand.Intn(40))

	// switchF()

	// if a == 32 && b == 41 {
	// 	println("32 is greater than 45")
	// } else if a == 45 || b == 40 {
	// 	println("a not equal 45 but b==40")
	// } else if a == 32 && a == 78 {
	// 	fmt.Println(" a equal to 32")
	// } else {
	// 	fmt.Println("No Value is true")
	// }

	if z := 2 * rand.Intn(b); z >= b {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v \n", z, b)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v \n", z, b)
	}

}
