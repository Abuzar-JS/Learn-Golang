package main

import "fmt"

// E59 exported to main.go
func E59() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foos(xi...))

	// xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(bars(xi))

}

func foos(ii ...int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func bars(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
