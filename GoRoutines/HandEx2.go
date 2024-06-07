package main

import "fmt"

type Person struct {
	first string
}

type Human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println("Hello")
}

func saysomething(h Human) {
	h.speak()
}

func HandEx2() {
	p1 := Person{
		first: "James",
	}

	saysomething(&p1)

	p1.speak()
}
