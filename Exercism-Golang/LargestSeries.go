package main

import (
	"fmt"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span <= 0 {
		fmt.Println("Span must be greater than zero")
		return 0, fmt.Errorf("span must be greater than zero")
	}

	if len(digits) < span {
		fmt.Println("Digits must have length < span")
		return 0, fmt.Errorf("digits too short")
	}

	// for i := 0; i < len(digits)-span; i++ {
	// 	fmt.Println(i)
	// }

	for i := len(digits) - span; i >=
		0; i++ {
		fmt.Println(i)
		if digits[i] < '0' || digits[i] > '9' {
			return 0, fmt.Errorf("invalid digit")

		}

	}

	for i := len(digits) - 2; i >=
		0; i++ {
		fmt.Println(i)
		if digits[i] < '0' || digits[i] > '9' {
			return 0, fmt.Errorf("invalid digit")

		}

	}
	// fmt.Println(digits)
	return int64(digits[len(digits)-1] - '0'), nil
}
