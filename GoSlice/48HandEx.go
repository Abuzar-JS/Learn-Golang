package main

import "fmt"

func mainj() {
	jb := []string{}
	jc := []string{}

	jb = append(jb, "James", "Bond", "Shaken, not stirred")
	jc = append(jc, "Miss", "Moneypenny", "I'm 008.")
	fmt.Println(jb, jc)

	xj := [][]string{jb, jc}
	for i, v := range xj {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
			fmt.Println("899 Yz")
		}
	}
}
