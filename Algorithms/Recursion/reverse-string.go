package main

import "fmt"

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseString(string(s[1:])) + string(s[0])
}

func main() {
	fmt.Println(reverseString("Hello, my name is Adam!"))
}
