package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to the calculator!\n" +
		"Enter an expression to evaluate or 'exit' to quit\n" +
		"Expression format: <number> <operation> <number>\n" +
		"Supported operations: +, -, *, /, //, %, **\n" +
		"Example: 1 + 2\n\n" +
		"The result of the operation '1 + 2' is: 3\n\n" +
		"Enter an expression:")

	for {
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		result := processExpression(strings.Split(text, " "))

		fmt.Printf("The result of the operation '%s' is: %f\n", text, result)
	}
}

func processExpression(tokens []string) float64 {
	var acc float64 = 0
	var operations []string = []string{"+", "-", "*", "/", "//", "%", "**"}
	var currentOp string

	for i, token := range tokens {
		if i%2 != 0 {
			if slices.Contains(operations, tokens[i]) {
				log.Fatalf("Invalid expression: '%s'\n", tokens)
			}
			currentOp = token
			continue
		}

		num, err := strconv.ParseFloat(token, 64)

		if err != nil {
			log.Fatal("Invalid number")
		}

		if i == 0 {
			acc = num
			continue
		}

		switch currentOp {
		case "+":
			acc += num
		case "-":
			acc -= num
		case "*":
			acc *= num
		case "/":
			acc /= num
		case "//":
			acc = float64(int(acc) / int(num))
		case "%":
			acc = float64(int(acc) % int(num))
		case "**":
			acc = float64(int(acc) ^ int(num))
		}
	}

	return acc
}
