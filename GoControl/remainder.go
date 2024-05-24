package main

import "fmt"

func mainRem() {
	x := 83 / 40
	y := 10 % 2
	fmt.Println(x, "\n", y)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("This is an even number \t %v \n", i)
		}
	}
}
