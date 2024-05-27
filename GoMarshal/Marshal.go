package main

import (
	"encoding/json"
	"fmt"
)

type Persons struct {
	Name       string
	Age        int
	Occupation string
}

func Mx() {
	p4 := Persons{
		Name:       "Arif",
		Age:        23,
		Occupation: "Developer",
	}
	p5 := Persons{
		Name:       "Aimen",
		Age:        26,
		Occupation: "Doctor",
	}
	people := []Persons{p4, p5}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
