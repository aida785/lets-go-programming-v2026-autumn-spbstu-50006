package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
	impossibleTemp = -1
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

func (tr *TemperatureRange) Apply(op Operator, value int) error {
	switch op {
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
		return fmt.Errorf("unknown operator: %s", op)
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
			return "", err
		}
		return "", fmt.Errorf("unexpected end of input")
	}
	return scanner.Text(), nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
	token, err := readToken(scanner)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(token)
}

func readOperator(scanner *bufio.Scanner) (Operator, error) {
	token, err := readToken(scanner)
	if err != nil {
		return "", err
	}

	op := Operator(token)
	switch op {
	case OpGTE, OpLTE, OpEQ:
		return op, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", op)
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

	for i := 0; i < departmentsCount; i++ {
		employeesCount, err := readInt(scanner)
		if err != nil {
			fmt.Println("Error reading employees count:", err)
			return
		}

		tempRange := NewTemperatureRange()

		for j := 0; j < employeesCount; j++ {
			op, err := readOperator(scanner)
			if err != nil {
				fmt.Println("Error reading operator:", err)
				return
			}

			val, err := readInt(scanner)
			if err != nil {
				fmt.Println("Error reading temperature:", err)
				return
			}

			if err := tempRange.Apply(op, val); err != nil {
				fmt.Println("Error applying temperature:", err)
				return
			}

			fmt.Println(tempRange.GetOptimal())
		}
	}
}
