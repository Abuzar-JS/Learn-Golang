package main

import "fmt"

func CallBack() {
	x := DoMath(10, 20, Addition)
	fmt.Println("Addition", x)

	x = DoMath(10, 20, Subtraction)
	fmt.Println("Subtraction", x)

}

func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func Addition(a int, b int) int {
	return a + b
}

func Subtraction(a int, b int) int {
	return a - b
}
