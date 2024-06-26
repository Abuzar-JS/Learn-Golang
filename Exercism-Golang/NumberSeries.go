package main

import (
	"fmt"
	"strconv"
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

	var product []string
	for i := 0; i <= len(digits)-span; i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return 0, fmt.Errorf("invalid digit")

		}
		product = append(product, digits[i:i+span])
	}
	//output[639 391 915]
	var largest int

	for _, num := range product {
		result := 1

		for _, char := range num {

			//Convert rune to string and then to int
			digit, err := strconv.Atoi(string(char))

			if err != nil {

				fmt.Println(err)
			}
			result *= digit
		}

		largest = getMax(largest, result)

	}

	return int64(largest), nil
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
