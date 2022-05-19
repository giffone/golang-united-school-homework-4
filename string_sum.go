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
	input = strings.Trim(input, " ")
	lInput := len(input)
	if lInput == 0 {
		return "", fmt.Errorf("%s", errorEmptyInput.Error())
	}
	buf := bytes.Buffer{}
	for index := 0; index < lInput; index++ {
		if input[index] == '-' {
			buf.WriteByte('+')
		}
		buf.WriteByte(input[index])
	}
	numS := strings.Split(buf.String(), "+")
	numN := []int{}
	for i := 0; i < len(numS); i++ {
		if numS[i] == "" {
			continue
		}
		n, err := strconv.Atoi(numS[i])
		if err != nil {
			return "", fmt.Errorf("%s", err.Error())
		}
		numN = append(numN, n)
	}
	lNumN := len(numN)
	if lNumN < 2 {
		return strconv.Itoa(numN[0]), nil
	}
	if lNumN > 2 {
		return "", fmt.Errorf("%s", errorNotTwoOperands.Error())
	}
	return strconv.Itoa(numN[0] + numN[1]), nil
}
