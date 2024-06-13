package main

import "fmt"

//PrintSquare exported
func PrintSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

//Square func
func Square(n int) int {
	return n * n
}
