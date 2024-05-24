package main

import "fmt"

func Basicmain() {
	city := map[string]int{
		"Abuzar": 42,
		"Ali":    21,
		"Kashif": 41,
		"Ahmed":  32,
	}
	fmt.Println("My name is Abuzar and i Live in city", city["Abuzar"], ".")

	fmt.Println(city)
	fmt.Printf("%#v\n", city)

	town := make(map[string]string)
	town["Abuzar"] = "Wapda"
	town["Ali"] = "Johar"
	town["Kashif"] = "TownShip"
	town["Ahmed"] = "Model"
	town["Uzair"] = "Model"

	fmt.Println(town["yalla"])
	fmt.Printf("My city %#v and town is %#v.", city["Abuzar"], town["Abuzar"])

	fmt.Println(len(town))
	fmt.Println(len(city))

}
