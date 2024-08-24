package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	// TODO: implement here
}

func decodeStringV1(s string) string {
	var stack []string
	for i := 0; i < len(s); i++ {
		word := string(s[i])
		if word != "]" {
			stack = append(stack, word)
			continue
		}

		// group words
		j := len(stack) - 1
		for ; j > 0; j-- {
			if top := stack[j]; top == "[" {
				break
			}
		}
		// +1 is `[`
		words := strings.Join(stack[j+1:], "")
		stack = stack[:j]

		// find number
		var count, base int
		j = len(stack)
		for ; j > 0; j-- {
			top := stack[j-1]
			if top != "0" && top != "1" && top != "2" && top != "3" && top != "4" &&
				top != "5" && top != "6" && top != "7" && top != "8" && top != "9" {
				break
			}
			number, _ := strconv.Atoi(top)
			count += int(math.Pow(10, float64(base)) * float64(number))
			base++
		}
		stack = stack[:j]
		stack = append(stack, strings.Repeat(words, count))
	}
	return strings.Join(stack, "")
}

// refer to: https://leetcode.com/problems/decode-string/solutions/4035120/go-solution-easy-explanation-o-n-time-space-complexity/?source=submission-ac
func decodeStringV2(s string) string {
	var stack []string
	for _, char := range s {
		word := string(char)
		if word != "]" {
			stack = append(stack, word)
			continue
		}

		// group words
		var words string
		for stack[len(stack)-1] != "[" {
			words = stack[len(stack)-1] + words
			stack = stack[:len(stack)-1]
		}
		// pop left `[`
		stack = stack[:len(stack)-1]

		// find number
		var number string
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top != "0" && top != "1" && top != "2" && top != "3" && top != "4" &&
				top != "5" && top != "6" && top != "7" && top != "8" && top != "9" {
				break
			}
			number = top + number
			stack = stack[:len(stack)-1]
		}
		count, _ := strconv.Atoi(number)
		stack = append(stack, strings.Repeat(words, count))
	}
	return strings.Join(stack, "")
}
