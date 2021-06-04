package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	slices := strings.Split(s, "")
	for i := len(slices)/2 - 1; i >= 0; i-- {
		opp := len(slices) - 1 - i
		slices[i], slices[opp] = slices[opp], slices[i]
	}
	return strings.Join(slices, "")
}

func main() {
	fmt.Printf("%s\n", reverse("Hello, my name is Adam!"))
}
