package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Enter a word: ")
	fmt.Scan(&s)
	s = strings.ToLower(s)
	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
