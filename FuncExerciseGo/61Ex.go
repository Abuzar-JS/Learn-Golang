package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v and age is %v", p.first, p.age)
}

func E61() {
	p1 := person{
		first: "Ali",
		age:   22,
	}
	person.speak(p1)
}
