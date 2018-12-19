package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string and I wil tell you if it starts with an 'i', ends with an 'a' and contains a 'n': ")
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSuffix(rawInput, "\n")

	if isTheFirstCharacterAnI(userInput) && isTheLastCharacterAnN(userInput) && isTheCharacterAAChracterOf(userInput) {
		fmt.Println("Found !")
	} else {
		fmt.Println("Not Found !")
	}

}

func isTheFirstCharacterAnI(userInput string) boolean {
	var firstCharacter = string(userInput[0])
	var output = (firstCharacter == "i" || firstCharacter == "I")
	return output
}

func isTheLastCharacterAnN(userInput string) boolean {
	var lastCharacter = string(userInput[len(userInput)-1 : 1])
	var output = (lastCharacter == "n" || lastCharacter == "N")
	return output
}

func isTheCharacterAAChracterOf(userInput string) boolean {
	var output = (strings.Contains(userInput, "a") || strings.Contains(userInput, "A"))
	return output
}
