package main

import "fmt"

func DelMap() {
	myMap := map[string]int{
		"a": 32,
		"b": 42,
		"c": 52,
		"d": 62,
		"e": 72,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	delete(myMap, "b")
	delete(myMap, "bs")

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	if a, ok := myMap["a"]; !ok {
		fmt.Println(a)
	} else {
		fmt.Println("Key did't exist")
	}
}
