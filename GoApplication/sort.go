package main

import (
	"fmt"
	"sort"
)

func SortGo() {

	xi := []int{1, 5, 477, 34, 98, 32, 6, 332, 986, 2132, 54, 76, 798, 23}
	xs := []string{"Animal", "Zebra", "Honey", "take x", "lets go", "Yx", "home", "Jenny", "Pakistan", "33Lah", "Rehaish", "My"}

	fmt.Println(xi)
	fmt.Println("------------------")
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("------------------")
	fmt.Println(xs)
	fmt.Println("------------------")
	sort.Strings(xs)
	fmt.Println(xs)
}
