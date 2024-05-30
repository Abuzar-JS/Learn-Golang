package main

import "fmt"

func Return70() func() int {
	return func() int { return 42 }
}

func E70() {

	a := Return70()
	fmt.Println(a())
}
