package main

import "fmt"

type people struct {
	firstname string
	lastname  string
	icecream  []string
}

func Main53() {
	people1 := people{
		firstname: "Ali",
		lastname:  "Ahmed",
		icecream:  []string{"Choclate", "Mango"},
	}

	people2 := people{
		firstname: "Saqib",
		lastname:  "Hameed",
		icecream:  []string{"Apple", "Vanila"},
	}

	fmt.Println("------------------")

	fmt.Println(people1)

	fmt.Println("------------------")
	fmt.Println(people2)

	fmt.Println("------------------")
	for i, v := range people2.icecream {
		fmt.Println(v, i)
	}
	fmt.Println("------------------")

	m := map[string]people{
		people1.lastname: people1,
		people2.lastname: people2,
	}

	fmt.Println("----------Hello--------")

	for _, v := range m {
		fmt.Println(v)
		fmt.Println("----------Break--------")
		for _, v2 := range v.icecream {
			fmt.Println(v.firstname, v.lastname, v2)
		}
	}
	fmt.Println("----------Hello--------")

}
