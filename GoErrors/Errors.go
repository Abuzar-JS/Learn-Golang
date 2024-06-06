package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Err1() {

	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

}

func ErAnswer() {
	var ans1, ans2, ans3 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		return
	}

	fmt.Println("Age: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		fmt.Println("This Error: ", err)
	}

	fmt.Println("Class: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		fmt.Println("This Error: ", err)
	}

	fmt.Println(&ans1)
	fmt.Println(ans1)
	fmt.Printf("%T", ans1)
}

func ErrFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Hey Abuzar")
	io.Copy(f, r)
}
