package main

import (
	"fmt"
)

func main() {
	var userInput float32
	fmt.Println("Enter a float:")
	fmt.Scanf("%f", &userInput)
	fmt.Println("This is your number", userInput)
	fmt.Println("This is your number truncated", trunc(userInput))
}

func trunc(floatNumber float32) int64 {
	return int64(floatNumber)
}
