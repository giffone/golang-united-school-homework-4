package string_sum

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
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
	lInput := len(input)
	if lInput == 0 {
		return "", fmt.Errorf("empty: %w", errorEmptyInput)
	}
	buf := bytes.Buffer{}
	for index := 0; index < lInput; index++ {
		if input[index] == ' ' {
			continue
		}
		if (input[index] == '-' || input[index] == '+') && index != 0 {
			buf.WriteByte(' ')
		}
		buf.WriteByte(input[index])
	}
	if buf.Len() == 0 {
		return "", fmt.Errorf("empty: %w", errorEmptyInput)
	}
	var number []int
	str := strings.Split(buf.String(), " ")
	for i := 0; i < len(str); i++ {
		if str[i] == "" || str[i] == "+" || str[i] == "-" {
			continue
		}
		n, err := strconv.Atoi(str[i])
		if err != nil {
			return "", fmt.Errorf("atoi: %w", err)
		}
		number = append(number, n)
	}
	if len(number) != 2 {
		return "", fmt.Errorf("more or less: %w", errorNotTwoOperands)
	}
	return strconv.Itoa(number[0] + number[1]), nil
}
