package main

import "fmt"

func Incre() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x += 33
		return x
	}
}
