package main

import "fmt"

// E57 exported to main.go
func E57() {
	x := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
	fmt.Printf("%T", x)

}

func sum(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
