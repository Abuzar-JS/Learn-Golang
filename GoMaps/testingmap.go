package main

import "fmt"

func TestingMap() {
	ice := make(map[string]int)
	ice["Hello"] = 3

	fmt.Println(ice["Hello"])
}
