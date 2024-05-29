package main

import (
	"bytes"
	"fmt"
)

func WriteInter() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello Gophers")

	// _, err = f.Write(s)

	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Its thursday")

	b.Write([]byte("Happ Happy"))
	fmt.Println(b.String())
}
