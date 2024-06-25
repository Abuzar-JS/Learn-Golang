package main

import (
	"fmt"
)

func mainZ() {
	numb := "639158"

	span := 3

	for i := 0; i < (len(numb) - span); i++ {
		fmt.Println(numb[i : i+3])
	}

	//for i := 0; i < len(numb); i++ {
	//	fmt.Println(numb[i : i+3])ss
	//}

	//for i := 0; i < (len(numb))+1; i++ {
	//	fmt.Println(numb[i])
	//}

}
