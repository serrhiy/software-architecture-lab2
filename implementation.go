package lab2

import (
	"errors"
)

const (
	NUMBER = iota
	PLUS
	MINUS
	MUL
	DIV
	POW
)

var operators = map[byte]int {
	'+': PLUS,
	'-': MINUS,
	'*': MUL,
	'/': DIV,
	'^': POW,
}

type Token struct {
	kind int
	source string
}

func isAlpha(char byte) bool {
	return char >= 48 && char <= 57
}

func tokenize(source string) ([]Token, error) {
	output := make([]Token, 0)
	for index := len(source) - 1; index >= 0; index-- {
		char := source[index]
		if (char == ' ') {
			continue
		}
		kind, isOperator := operators[char]
		if isOperator {
			output = append(output, Token{ kind, string(char) })
			continue
		}
		number := ""
		for char := source[index]; isAlpha(char); {
			number = string(char) + number
			index -= 1
			if index < 0 {
				break
			}
			char = source[index]
		}
		if (len(number) == 0) {
			return nil, errors.New("Invalid character: " + string(source[index]))
		}
		index += 1;
		output = append(output, Token{ NUMBER, number })
	}
	return output, nil
}

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToInfix(source string) (string, error) {
	tokens, validSource := tokenize(source)
	if validSource != nil {
		return "", validSource
	}
	stack := make([]string, 0)
	for _, token := range tokens {
		if token.kind == NUMBER {
			stack = append(stack, token.source)
			continue
		}
		length := len(stack)
		operand1, operand2 := stack[length - 1], stack[length - 2]
		expression := operand1 + " " + token.source + " " + operand2
		stack[length - 2] = expression
		stack = stack[: length - 1]
	}
	return stack[0], nil
}
