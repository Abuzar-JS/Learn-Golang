package main

import "fmt"

func mainx() {
	x := 40
	switch x {
	case 40:
		fmt.Println(" x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("Printed because of all of the fallthrough statements: this is the default case for x")
	}
}
