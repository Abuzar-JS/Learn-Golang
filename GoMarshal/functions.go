package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"customName"`
	Age  int    `json:"age,omitempty"`
}

func JsonU() {
	p := Person{Name: "tom"}
	pBytes, err := json.Marshal(p)

	log.Print(err)
	log.Print(string(pBytes))

}
