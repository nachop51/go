package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type myString string

func (s myString) Split(sep string) []string {
	var tokens []string
	var token string

	// "   1  +    2   " -> ["1", "+", "2"]

	var sLen, sepLen int = len(s), len(sep)
	var i int = 0

	for i < sLen {
		for i+sepLen < sLen && s[i:i+sepLen] == myString(sep) {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			i += sepLen
		}

		token += string(s[i])
		i++
	}

	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens
}

func (s myString) Strip() myString {
	var start, end int

	for i, char := range s {
		if char != ' ' && char != '\n' {
			start = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && s[i] != '\n' {
			end = i
			break
		}
	}

	return myString(s[start : end+1])
}

func (s myString) Compare(other myString) bool {
	return s == other
}

func contains(arr []string, item string) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}

	return false
}

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

		myStr := myString(text).Strip()
		// fmt.Println(myStr)

		if myStr.Compare("exit") {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		result := processExpression(myStr.Split(" "))

		fmt.Printf("The result of the operation '%s' is: %f\n", myStr, result)
	}
}

func printTokens(tokens []string) {
	for _, token := range tokens {
		fmt.Printf("|%s|\n", token)
	}
}

func processExpression(tokens []string) float64 {
	var acc float64 = 0
	var operations []string = []string{"+", "-", "*", "/", "//", "%", "**"}
	var currentOp string

	for i, token := range tokens {
		if i%2 != 0 {
			if !contains(operations, tokens[i]) {
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
