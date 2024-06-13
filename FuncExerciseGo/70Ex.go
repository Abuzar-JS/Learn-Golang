package main

import "fmt"

//Return70 exported
func Return70() func() int {
	return func() int { return 42 }
}

//E70 exported
func E70() {

	a := Return70()
	fmt.Println(a())
}
