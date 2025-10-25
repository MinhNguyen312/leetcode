package main

import "fmt"

func isValid(s string) bool {

	if len(s)%2 != 0 || s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}

	// check for valid
	// valid case
	var parentheses []rune

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			parentheses = append(parentheses, c)
			continue
		}

		switch c {
		case ')':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '(' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		case ']':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '[' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		case '}':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '{' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		default:
			return false
		}
	}

	if len(parentheses) != 0 {
		return false
	}

	return true
}

func main() {
	tests := []string{
		"(){}[]", "()", "(]", "({[]})", "([)]",
	}

	for _, test := range tests {
		fmt.Printf("%s is valid: %v\n", test, isValid(test))
	}

}
