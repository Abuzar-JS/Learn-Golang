package main

import "fmt"

type Person struct {
	first string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (p Person) speak() {
	fmt.Println("I am ", p.first)

}
func (sa SecretAgent) speak() {
	fmt.Println("I am a secret agent ", sa.first)

}

type human interface {
	speak()
}

func SaySomething(h human) {
	h.speak()
}

func MethodX() {
	sa1 := SecretAgent{
		Person: Person{first: "James"},
		ltk:    true,
	}
	// p1 := Person{
	// 	first: "James",
	// }
	p2 := Person{
		first: "Jenny",
	}
	p3 := Person{
		first: "Scuba",
	}

	sa1.speak()
	p2.speak()
	p3.speak()

	fmt.Println("----------")
	SaySomething(sa1)
	SaySomething(p2)
}
