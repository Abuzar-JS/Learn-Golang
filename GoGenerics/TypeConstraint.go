package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | string | float64](a, b T) T {
	return a + b

}

type myNumbers interface {
	int | float64 | string
}

func addMyNumbers[T myNumbers](a, b T) T {
	return a + b
}


func TConst() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2))
	fmt.Println(addT("3333", "333"))
	fmt.Println(addMyNumbers("H", "ds"))

}
