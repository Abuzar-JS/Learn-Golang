package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HasP is exported to main.go
func HasP() {
	s := "12345pass"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)

	}

	loginpass := "12345pass"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpass))
	if err != nil {
		fmt.Println("Wrong Pass")
		return
	}
	fmt.Println("Password Matched")

}
