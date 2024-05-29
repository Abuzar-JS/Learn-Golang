package main

import "fmt"

func E58() {

	a := foo()
	fmt.Println(a)
	c, d := bar()
	fmt.Println(c, d)
}

func foo() int {
	return 29
}

func bar() (int, string) {
	return 22, "Abuzar"
}
