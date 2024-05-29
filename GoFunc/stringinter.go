package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
}

func (b Book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func StX() {
	b := Book{
		title: "West With The Night",
	}
	var n count = 42
	fmt.Println(b)
	fmt.Println(n)
}
