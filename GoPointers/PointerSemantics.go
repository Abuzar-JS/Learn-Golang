package main

import "fmt"

func addone(i int) int {
	return i + 1
}

func addoneP(v *int) {
	*v += 1
}

func PoSematic() {

	//Value Semantics
	a := 1
	fmt.Println(a)
	fmt.Println(addone(a))
	fmt.Println(a)

	// Pointer Semantics
	b := 1
	fmt.Println("Value Semantics")
	fmt.Println(b)
	addoneP(&b)
	fmt.Println(b)

}
