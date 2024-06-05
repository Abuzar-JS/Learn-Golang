package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name    string
	Age     int
	Class   string
	School  string
	Teacher []teacher
}

type teacher struct {
	Id      int
	Name    string
	Subject string
}

func MarshalJ() {

	s1 := student{
		Name:   "Ali",
		Age:    10,
		Class:  "2b",
		School: "The Educators",
		Teacher: []teacher{{Id: 1,
			Name:    "Sir Awais",
			Subject: "English"},
		}}

	s2 := student{
		Name:   "Ahmed",
		Age:    12,
		Class:  "4C",
		School: "The Marshal Boys High School",
		Teacher: []teacher{{Id: 3,
			Name:    "Sir Shakir",
			Subject: "Chemistry"}}}

	Sd := []student{s1, s2}
	fmt.Println(Sd)
	Jd, err := json.Marshal(Sd)
	if err != nil {
		return
	}
	fmt.Println(Jd)
	fmt.Println(string(Jd))

}
