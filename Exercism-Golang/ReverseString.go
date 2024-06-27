package main

import "fmt"

func Reverse(input string) string {

	reverse := []rune(input)

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	fmt.Println(string(reverse))
	return string(reverse)
}
