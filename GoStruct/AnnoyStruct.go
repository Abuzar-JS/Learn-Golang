package main

import "fmt"

type pers struct {
	name     string
	age      int
	location string
}

func AnnoyStruct() {
	p1 := pers{
		// 	name     string
		// 	age      int
		// 	location string
		// }{
		name:     "ali",
		age:      34,
		location: "lahore",
	}

	fmt.Println(p1.location)
}
