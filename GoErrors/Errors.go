package main

import (
	"fmt"
	"io"
	"log"
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

	r := strings.NewReader("Hey Abuzar1 \n")
	io.Copy(f, r)

	x := strings.NewReader("Hey Abuzar2 \n")
	io.Copy(f, x)

	z := strings.NewReader("Hey Abuzar3 \n")
	io.Copy(f, z)
}

func ErrOpenF() {
	f, err := os.Open("names.taxt")
	if err != nil {
		log.Println("Log", err)
		fmt.Println("Log", err)
		// fmt.Println(err)
		return
		defer f.Close()
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
