package main

func Reverse(input string) string {

	reverse := []rune(input)

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return string(reverse)
	//var test string
	//
	//for i := len(input) - 1; i >= 0; i-- {
	//	test = test + string(input[i])
	//
	//}
	//
	//fmt.Println(test)
	//
	//return test
}
