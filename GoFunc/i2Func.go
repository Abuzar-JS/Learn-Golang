package main

import (
	"fmt"
)

// Func with no parameter no return
func Foos() {
	fmt.Println("I'm Abuzar")
}

// func with 1 para but no returns
func Bar(a string) {
	fmt.Println("My name is ", a)

}

// func with 1 para 1 retunr
func Aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

// func with 2 para 2 returns

func Dogyears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is old in dog years"), age
}

// func (receiver) identifier(parameteres)(returns ){code}
