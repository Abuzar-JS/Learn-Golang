package main

import "fmt"

//AnnoyX is exported
func AnnoyX() {
	func() {
		for i := 0; i < 80; i++ {
			fmt.Println(i)
		}
	}()
}

//a
