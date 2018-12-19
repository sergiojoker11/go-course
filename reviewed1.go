package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter floating point number: ")

	rawInput, _ := reader.ReadString('\n')

	input := strings.TrimSuffix(rawInput, "\n")

	inputToFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Please enter valid floating point number: ")
		return
	}

	resultOutput := int(inputToFloat)

	fmt.Println(resultOutput)

}
