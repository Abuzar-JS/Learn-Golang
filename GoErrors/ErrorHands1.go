package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func ErrH1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, stirred", "Any last wishes?", "Never say never"},
	}

	fmt.Println(p1)

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Json Has Errors", err)
		return
	}
	fmt.Println(string(bs))

}
