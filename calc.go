package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrorDivisionByZero        = errors.New("division by zero")
	ErrorInvalidExpression     = errors.New("invalid expression")
	ErrorMismatchedParentheses = errors.New("mismatched parentheses")
	ErrorInvalidCharacter      = errors.New("invalid character")
)

// Calc evaluates a mathematical expression given as a string.
func Calc(expression string) (float64, error) {
	tokens := extract_tokens(expression)
	postfix, err := to_postfix(tokens)
	if err != nil {
		return 0, err
	}
	return eval(postfix)
}

func extract_tokens(expr string) []string {
	var tokens []string
	var current_token strings.Builder

	for _, char := range expr {
		switch char {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if current_token.Len() > 0 {
				tokens = append(tokens, current_token.String())
				current_token.Reset()
			}
			tokens = append(tokens, string(char))
		default:
			current_token.WriteRune(char)
		}
	}

	if current_token.Len() > 0 {
		tokens = append(tokens, current_token.String())
	}

	return tokens
}

func to_postfix(tokens []string) ([]string, error) {
	var output []string
	var operators []string

	for _, token := range tokens {
		if is_number(token) {
			output = append(output, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, ErrorMismatchedParentheses
			}
			operators = operators[:len(operators)-1] // Pop the '('
		} else if is_operator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return nil, ErrorInvalidCharacter
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, ErrorMismatchedParentheses
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func eval(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if is_number(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if is_operator(token) {
			if len(stack) < 2 {
				return 0, ErrorInvalidExpression
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, ErrorDivisionByZero
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		} else {
			return 0, fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return 0, ErrorInvalidExpression
	}

	return stack[0], nil
}

func is_number(token string) bool {
	if _, err := strconv.ParseFloat(token, 64); err == nil {
		return true
	}
	return false
}

func is_operator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

// func main() {
// 	expression := "3 + 5 * (2 - 8)"
// 	result, err := Calc(expression)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }
