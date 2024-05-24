package main

import "fmt"

func CityStruc() {
	type location struct {
		name    string
		temp    int
		state   string
		Country string
	}

	city1 := location{
		name:    "Lahore",
		temp:    45,
		state:   "Punjab",
		Country: "Pakistan",
	}

	city2 := location{
		name:    "Peshawar",
		temp:    38,
		state:   "Khyber Pakhtunkhwa",
		Country: "Pakistan",
	}
	city3 := location{
		name:    "Karachi",
		temp:    32,
		state:   "Sindh",
		Country: "Pakistan",
	}

	fmt.Printf("%v \n %v \n %v", city1, city2, city3)

	fmt.Println("--------------------------------")
	fmt.Println(city1.name)
	fmt.Println("--------------------------------")
	fmt.Println(city3.temp)
	fmt.Println("--------------------------------")
	fmt.Println(city2.name)
	fmt.Println("--------------------------------")
	fmt.Println(city1.state)
	fmt.Println("--------------------------------")
	fmt.Println(city2.Country)
	fmt.Println("--------------------------------")
}
