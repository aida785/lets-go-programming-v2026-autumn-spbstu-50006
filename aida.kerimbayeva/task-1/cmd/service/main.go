package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, op string

	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	x, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(x / y)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
