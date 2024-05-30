package main

import "fmt"

func intDelta(n *string) {
	*n = "Hello abx"
}

func sliceDelta(ii []int) {
	ii[0] = 99
}
func MapDelta(md map[string]int, s string) {
	md[s] = 33
}

func ValuePass() {
	a := 42
	x := "Abuzar"
	fmt.Println(x)
	intDelta(&x)
	fmt.Println(x)
	fmt.Println(a)

	c := 82
	fmt.Println(c)
	// intDelta(&c)
	fmt.Println(c)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["james"] = 42
	fmt.Println(m["James"])
	MapDelta(m, "James")
	fmt.Println(m["James"])
}
