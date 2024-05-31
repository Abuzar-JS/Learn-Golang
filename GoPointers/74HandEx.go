package main

import "fmt"

func E74() {
	a := 3
	fmt.Println(a)
	fmt.Println(&a)

	for i := 0; i <= 200; i++ {
		fmt.Printf("Value of A %v \t%v and \t Loop Range %v \n", a, &a, i)
		a += 1

	}
}
