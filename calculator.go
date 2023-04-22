package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (c calc) operate(input string, operation string) (int, error) {

	cleanInput := strings.Split(input, operation)

	first, err := c.parseString(cleanInput[0])

	if err != nil {
		return 0, err
	}

	second, err := c.parseString(cleanInput[1])

	if err != nil {
		return 0, err
	}

	switch operation {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		return first / second, nil
	default:
		fmt.Println("Invalid operator")
		return 0, nil
	}
}

func (calc) parseString(operator string) (int, error) {
	result, _ := strconv.Atoi(operator)
	return result, nil
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Enter your input:")
	input := readInput()
	fmt.Println("Enter your operation:")
	operator := readInput()

	processResult(input, operator)
}

func processResult(input string, operator string) {
	c := calc{}
	value, err := c.operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of:", input, "equals to", value)
	}
}
