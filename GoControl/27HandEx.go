package main

import "fmt"

func main27() {
	x := 20
	for {
		if x < -20 {
			break
		}
		fmt.Println(x)
		x--
	}
}
