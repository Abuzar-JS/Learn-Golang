package main

import (
	"fmt"
	"log"
	"os"
)

func WrapMax() {
	xb, err := readfile("poem.txt")
	if err != nil {
		log.Fatalf("readfile in main: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readfile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil,
			fmt.Errorf("error in readfile: %s", err)
	}
	return xb, nil
}
