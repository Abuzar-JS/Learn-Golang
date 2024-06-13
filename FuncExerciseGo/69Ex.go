package main

import "fmt"

//E69 exported
func E69() {
	y := func() {
		for i := 0; i < 80; i++ {
			fmt.Println(i)
		}
	}
	y()
	x()
}

var x = func() {
	for i := 0; i < 80; i++ {
		fmt.Println(i)
	}
}
