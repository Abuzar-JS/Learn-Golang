package main

import "fmt"

func loopX() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Counting to five: %v \t first for loop \n", i)
	// }

	// var y int = 5

	// for y < 10 {
	// 	fmt.Printf("y is %v \t\t\t second for loop \n", y)
	// 	y++
	// }

	// for {
	// 	fmt.Printf("y is %v \t\t third for loop \n", y)
	// 	if y == 200 {
	// 		break
	// 	}
	// 	y += 3

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Counting even numbers: ", i)
	}

}
