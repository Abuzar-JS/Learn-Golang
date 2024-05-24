package main

import "fmt"

func mainS() {
	a := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)"}

	// fmt.Println(a)
	// for i, v := range a {
	// 	fmt.Printf("%v - value %v \n", i, v)

	// for _, v := range a {
	// 	fmt.Printf(" %v \n", v)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		fmt.Println("Value of i =", i)
		fmt.Println("-----------")
	}
}
