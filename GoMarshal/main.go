package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name       string `json:"customName"`
	Age        int    `json:"age,omitempty"`
	CreditCard string `json:"-"`
}

func main() {
	p := Person{Name: "tom", CreditCard: "super secret"}
	pBytes, err := json.Marshal(p)

	log.Print(err)
	log.Print(string(pBytes))

	p2AsRawJson := json.RawMessage(`{"customName":"mary"}`)
	var p2 Person
	err2 := json.Unmarshal(p2AsRawJson, &p2)
	log.Print(err2)
	log.Printf("%+v", p2)

}
