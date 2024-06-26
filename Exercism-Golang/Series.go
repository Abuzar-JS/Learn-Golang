package main

func All(n int, s string) []string {

	var substring []string

	for i := 0; i <= len(s)-n; i++ {
		substring = append(substring, s[i:i+n])
	}
	return substring
}

func UnsafeFirst(n int, s string) string {
	var substring []string

	for i := 0; i <= len(s)-n; i++ {
		substring = append(substring, s[i:i+n])
	}
	return substring[0]
}
