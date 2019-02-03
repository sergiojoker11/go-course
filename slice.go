package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	userNumberStorage := make([]int, 3)
	var userInput string
	for ok := true; ok; ok = !isTheExitCharacter(userInput) {
		userInput = askForUserInput()
		userNumberInput, err := strconv.Atoi(userInput)
		if err != nil || !isTheExitCharacter(userInput) {
			userNumberStorage = append(userNumberStorage, userNumberInput)
			sortStorage(userNumberStorage)
			print(userNumberStorage)
		}
	}
}

func isTheExitCharacter(userInput string) bool {
	var firstCharacter = string(userInput[0])
	var output = (firstCharacter == "x" || firstCharacter == "X")
	return output
}

func askForUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integers followed by ENTER or enter 'x' or 'X' to exit:")
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSuffix(rawInput, "\n")
	return userInput
}

func sortStorage(userNumberStorage []int) {
	sort.Ints(userNumberStorage)
}

func print(userNumberStorage []int) {
	fmt.Println(userNumberStorage)
}
