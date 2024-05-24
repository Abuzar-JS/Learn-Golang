package main

import "fmt"

type amst struct {
	first     string
	friends   int
	favDrinks []string
}

func Main56() {

	elif := struct {
		first     string
		friends   int
		favDrinks []string
	}{
		first:     "Ali",
		friends:   4,
		favDrinks: []string{"Cola", "Fanta"},
	}
	plif := struct {
		first     string
		friends   int
		favDrinks []string
	}{
		first:     "Pervaiz",
		friends:   2,
		favDrinks: []string{"sprite", "Mint"},
	}

	fmt.Println(elif.first, elif.friends, elif.favDrinks)
	fmt.Println("-------------------------")
	fmt.Println(plif.first, plif.friends, plif.favDrinks)

}
