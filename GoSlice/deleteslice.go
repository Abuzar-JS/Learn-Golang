package main

import "fmt"

func main() {
	x2 := []int{22, 33, 43, 53, 63, 73, 83, 93}
	fmt.Printf("x2 - %v \n", x2)
	fmt.Println("-----------")

	x2 = append(x2[:3], x2[5:]...)
	fmt.Println(x2)
}
