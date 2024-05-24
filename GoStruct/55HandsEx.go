package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func Main55() {
	vehicle1 := vehicle{
		engine: engine{false},
		make:   "Toyota",
		model:  "Corolla",
		doors:  4,
		color:  "white",
	}

	vehicle2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "Model Y",
		doors: 2,
		color: "Black",
	}

	fmt.Println(vehicle1)
	fmt.Println("-----------------")
	fmt.Println(vehicle1.engine, vehicle1.color, vehicle1.make, vehicle1.model)
	fmt.Println("-----------------")
	fmt.Println(vehicle2)
	fmt.Println("-----------------")
	fmt.Println(vehicle2.engine, vehicle2.color, vehicle2.make, vehicle2.model)

}
