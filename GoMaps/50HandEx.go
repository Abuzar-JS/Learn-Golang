package main

import "fmt"

func Main50() {

	// bond := []string{"shaken, not stirred", "martinis", "fast cars"}

	// fmt.Println(bond)

	myMap := make(map[string][]string)

	myMap["Jack"] = []string{"Hello Abuzar", "Hello Jazz"}

	myMap["bond_james"] = append(myMap["bond_james"], "shaken, not stirred", "martinis", "fast cars")
	myMap["money_penny"] = append(myMap["money_penny"], "intelligence", "literature", "computer science")
	myMap["no_doctor"] = append(myMap["no_doctor"], "cats", "ice cream", "sunset")
	myMap["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	// fmt.Println("\n", myMap["bond_james"])

	for k, v := range myMap {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

	fmt.Println("---------------")

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
