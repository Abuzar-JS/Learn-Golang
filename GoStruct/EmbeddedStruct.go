package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person
	cnic int
}

func EmbStruc() {
	emp1 := employee{
		person: person{
			name: "ali",
			age:  20},
		cnic: 1001,
	}

	emp2 := employee{
		person: person{name: "Saqib", age: 32},
		cnic:   1002,
	}
	fmt.Println(emp1.age)
	fmt.Println("-----------------")
	fmt.Println(emp1.cnic)
	fmt.Println("-----------------")
	fmt.Println(emp2.name)
	fmt.Println("-----------------")
	fmt.Println(emp2.person)
	fmt.Println("-----------------")

	fmt.Println(emp1)
	fmt.Println("-----------------")
	fmt.Println("-----------------")

	fmt.Println(emp2)

	fmt.Printf("%T \t %#v\n", emp1, emp1)
	fmt.Println("-----------------")

	fmt.Println(emp1.name, emp1.age, emp1.age)
	fmt.Println("-----------------")
	fmt.Println(emp1.name, emp1.age, emp1.age, emp1.cnic, emp1.person)
	fmt.Println("-----------------")
	fmt.Println(emp1.name, emp1.person.name)
}
