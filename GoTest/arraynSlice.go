package main

import (
	"fmt"
)

func SliceF() {
	c := [3]int{1, 2, 3}
	b := [...]string{"Hello", "Abuzar"}
	d := []int{19, 66, 44, 22, 77, 88}

	fmt.Printf("len of c is %#v and all elements are %v \n", len(c), c)
	fmt.Printf("len of b is %#v and all elements are %v and type is %T \n", len(b), b, b)
	fmt.Printf("len of c is %#v and index at '2' is %v \n", cap(c), c[2])
	fmt.Printf("len of d is %#v and all elements are %v \n", len(d), d)
	fmt.Printf("on index [0] of b is: %v \n", b[0])
	fmt.Println(d[1:5])

	d = append(d, 888, 9999, 555)
	fmt.Println("--------------------")
	fmt.Println(d)
	d = append(d[:2], d[5:]...)
	y := make([]int, len(d))
	copy(y, d)
	fmt.Println(y)
	fmt.Println(d)

	cd = [][]int[c, d]
	fmt.Println(cd)

}
