package main

import "fmt"

func mainAP() {
	xi := []int{22, 33, 43, 53, 63, 73, 83, 93}
	// fmt.Println(xi)
	// xi = append(xi, 92, 99, 87, 44, 56, 47, 28)

	// fmt.Println(xi)

	// Slicing the slice
	//[0:3] [inclusive:Exclusive]

	fmt.Println(xi[3:])  //Result 53,63,73,83,93
	fmt.Println(xi[:6])  //Result 22, 33, 43, 53, 63, 73
	fmt.Println(xi[2:7]) //Result  22, 33, 43, 53, 63, 73, 83,
	fmt.Println(xi[2:7]) //Result  22, 33, 43, 53, 63, 73, 83, 93
	fmt.Println(xi[5:8]) //Result   73, 83, 93

}
