package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
	impossibleTemp = -1
)

var (
	ErrUnknownOperator      = errors.New("unknown operator")
	ErrUnexpectedEndOfInput = errors.New("unexpected end of input")
)

type Operator string

const (
	OpGTE Operator = ">="
	OpLTE Operator = "<="
	OpEQ  Operator = "="
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		min: minAllowedTemp,
		max: maxAllowedTemp,
	}
}

func (tr *TemperatureRange) Apply(operator Operator, value int) error {
	switch operator {
	case OpGTE:
		if value > tr.min {
			tr.min = value
		}
	case OpLTE:
		if value < tr.max {
			tr.max = value
		}
	case OpEQ:
		if value > tr.min {
			tr.min = value
		}

		if value < tr.max {
			tr.max = value
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnknownOperator, operator)
	}

	return nil
}

func (tr *TemperatureRange) GetOptimal() int {
	if tr.min > tr.max {
		return impossibleTemp
	}

	return tr.min
}

func readToken(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("scanner error: %w", err)
		}

		return "", ErrUnexpectedEndOfInput
	}

	return scanner.Text(), nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
	token, err := readToken(scanner)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(token)
	if err != nil {
		return 0, fmt.Errorf("parse integer: %w", err)
	}

	return value, nil
}

func readOperator(scanner *bufio.Scanner) (Operator, error) {
	token, err := readToken(scanner)
	if err != nil {
		return "", err
	}

	operator := Operator(token)
	switch operator {
	case OpGTE, OpLTE, OpEQ:
		return operator, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnknownOperator, operator)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	departmentsCount, err := readInt(scanner)
	if err != nil {
		fmt.Println("Error reading departments count:", err)
		return
	}

	for range departmentsCount {
		employeesCount, err := readInt(scanner)
		if err != nil {
			fmt.Println("Error reading employees count:", err)
			return
		}

		tempRange := NewTemperatureRange()

		for range employeesCount {
			operator, err := readOperator(scanner)
			if err != nil {
				fmt.Println("Error reading operator:", err)
				return
			}

			val, err := readInt(scanner)
			if err != nil {
				fmt.Println("Error reading temperature:", err)
				return
			}

			if err := tempRange.Apply(operator, val); err != nil {
				fmt.Println("Error applying temperature:", err)
				return
			}

			fmt.Println(tempRange.GetOptimal())
		}
	}
}