package main

import "fmt"

func Sum(ii []int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func test(){
	arr := []int{1,2,3,4}
	Sum(arr)
}
