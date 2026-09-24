package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, op string

	fmt.Scanln(&a)
	x, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	fmt.Scanln(&b)
	y, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Scanln(&op)

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
