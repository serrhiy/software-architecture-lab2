package lab2

import (
	"errors"
	"strconv"
	"strings"
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

type Component struct {
	source string
	isComposite bool
}

func isAlpha(char byte) bool {
	return char >= 48 && char <= 57
}

func isValidInteger(number string) bool {
	empty := len(number) == 0
	startsWith0 := strings.HasPrefix(number, "0")
	startsWithMinus0 := strings.HasPrefix(number, "-0")
	startsWithPlus0 := strings.HasPrefix(number, "+0")
	if empty || startsWith0 || startsWithMinus0 || startsWithPlus0 {
		return false
	}
	_, err := strconv.Atoi(number)
	return err == nil
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
		if len(number) == 0 {
			return nil, errors.New("Invalid character: " + string(source[index]))
		}
		if !isValidInteger(number) {
			return nil, errors.New("Invalid number: " + number)
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
	stack := make([]Component, 0)
	for _, token := range tokens {
		if token.kind == NUMBER {
			stack = append(stack, Component{ token.source, false })
			continue
		}
		length := len(stack)
		operand1, operand2 := stack[length - 1], stack[length - 2]
		operand1Str, operand2Str := operand1.source, operand2.source
		if token.kind == MUL || token.kind == DIV {
			if operand1.isComposite {
				operand1Str = "(" + operand1Str + ")"
			}
			if operand2.isComposite {
				operand2Str = "(" + operand2Str + ")"
			}
		}
		expression := operand1Str + " " + token.source + " " + operand2Str
		stack[length - 2] = Component{ expression, true }
		stack = stack[: length - 1]
	}
	return stack[0].source, nil
}
