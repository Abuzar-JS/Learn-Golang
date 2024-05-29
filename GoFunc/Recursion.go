package main

import "fmt"

func FacMain() {
	factori(4)
	fmt.Println("--------------")
	fmt.Println(FactLopp(10))

}

func factori(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1

	}
	return n * factori(n-1)

}

func FactLopp(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total

}
