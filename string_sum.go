package string_sum

// package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")

	errorIsNotNumber = errors.New("Operand is not number")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var operandAToken []string
	var operandBToken []string

	nextOperator := false

	prev := ""
	for _, v := range input {
		value := strings.Trim(string(v), " ")
		if value == "" {
			continue
		}

		if nextOperator == false {
			if prev != "null" {
				if isOperator(value) && isNumber(prev) {
					nextOperator = true
				}
			}
		}

		if nextOperator == false {
			operandAToken = append(operandAToken, value)
		} else {
			if isNumber(prev) && isOperator(value) && len(operandBToken) > 0 {
				return "", fmt.Errorf("%w", errorNotTwoOperands)
			}

			operandBToken = append(operandBToken, value)
		}

		prev = value
	}

	return getFromOperandTokens(operandAToken, operandBToken)
}

func isNumber(value string) bool {
	r := []rune(value)

	if len(r) == 0 {
		return false
	}
	return r[0] > '0' && r[0] < '9'
}

func isOperator(value string) bool {
	return value == "+" || value == "-"
}

func getFromOperandTokens(opA []string, opB []string) (str string, err error) {
	as := strings.Join(opA, "")
	bs := strings.Join(opB, "")

	if len(as) == 0 || len(bs) == 0 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	a, errA := strconv.Atoi(as)
	b, errB := strconv.Atoi(bs)

	if a == 0 && b == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	if errA != nil {
		return "", fmt.Errorf("%w", errA)
	}

	if errB != nil {
		return "", fmt.Errorf("%w", errB)
	}

	return strconv.Itoa(a + b), nil
}

// func main() {
// 	fmt.Println(StringSum("24c + 55"))
// }
