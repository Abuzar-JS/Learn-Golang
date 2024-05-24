package main

import "fmt"

func Rangemap() {

	an := make(map[string]int)

	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v \n", an)
	fmt.Println(len(an))

	fmt.Println("--------0---------")

	for k, v := range an {
		fmt.Println("\n", k, v)
	}

	fmt.Println("--------1---------")

	for _, v := range an {
		fmt.Println("\n", v)
	}

	fmt.Println("-------2----------")

	for k := range an {
		fmt.Println("\n", k)
	}

	// for range with a slice

	fmt.Println("-----------------")

	xi := []int{42, 43, 44}

	for k, v := range xi {
		fmt.Println("\n", k, v)
	}

	fmt.Println("-------1---------")

	for _, v := range xi {
		fmt.Println("\n", v)
	}

	fmt.Println("---------2--------")

	for k := range xi {
		fmt.Println("\n", k)
	}

}
