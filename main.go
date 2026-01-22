package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getUserInput(r io.Reader) float64 {

	reader := bufio.NewReader(r)
	fmt.Println("Please insert the temperature: ")

	for {

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		input = strings.TrimSpace(input)
		userNum, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Sorry I couldn't convert that input")
			continue
		}
		return userNum
	}
}

func getUserUnit(r io.Reader) string {
	reader := bufio.NewReader(r)
	fmt.Println("Please insert the unit of the temperature: ")

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input!")
			continue
		}
		input = strings.TrimSpace(input)

		if len(input) > 0 {
			Unit := string(input[0])
			if Unit == "C" || Unit == "F" {
				return Unit
			}
			fmt.Println("Invalid unit, please use F or C")
		}
	}
}

func cToF(temp float64) float64 {
	return (temp * 1.8) + 32
}

func fToC(temp float64) float64 {
	return (temp - 32) / 1.8
}

func main() {
	var result float64

	temp := getUserInput(os.Stdin)
	unit := getUserUnit(os.Stdin)

	switch unit {
	case "c", "C":
		result = cToF(temp)
		fmt.Println(result)

	case "f", "F":
		result = fToC(temp)
		fmt.Println(result)

	}
}
