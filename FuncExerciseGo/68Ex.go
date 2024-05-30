package main

import "fmt"

func AnnoyX() {
	func() {
		for i := 0; i < 80; i++ {
			fmt.Println(i)
		}
	}()
}
