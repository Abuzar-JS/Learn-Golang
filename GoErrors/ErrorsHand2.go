package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First   string
	Last    string
	Sayings []string
}

func ErrH2() {
	p1 := Person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error marshalling %v in toJSON: %v", a, err)
	}
	return bs, nil
}
