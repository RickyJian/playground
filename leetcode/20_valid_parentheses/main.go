package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("{{"))

	fmt.Println(isValidEnhance("{}"))
	fmt.Println(isValidEnhance("()[]{}"))
	fmt.Println(isValidEnhance("{{"))
}

func isValid(s string) bool {
	var left = map[rune]struct{}{
		'(': {},
		'[': {},
		'{': {},
	}
	var right = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var parentheses []rune
	for _, word := range s {
		if _, ok := left[word]; ok {
			parentheses = append(parentheses, word)
		} else if mapping, ok := right[word]; ok {
			if len(parentheses) == 0 || mapping != parentheses[len(parentheses)-1] {
				return false
			}
			parentheses = parentheses[:len(parentheses)-1]
		}
	}
	return len(parentheses) == 0
}

func isValidEnhance(s string) bool {
	parentheses := make([]rune, 0, len(s)/2)
	for _, word := range s {
		switch word {
		case '(':
			parentheses = append(parentheses, ')')
		case '[':
			parentheses = append(parentheses, ']')
		case '{':
			parentheses = append(parentheses, '}')
		default:
			lastIdx := len(parentheses) - 1
			if len(parentheses) == 0 || word != parentheses[lastIdx] {
				return false
			}
			parentheses = parentheses[:lastIdx]
		}
	}
	return len(parentheses) == 0
}
