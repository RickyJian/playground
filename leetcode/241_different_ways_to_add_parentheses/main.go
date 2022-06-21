package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

func diffWaysToCompute(expression string) []int {
	var results []int
	for i, exp := range expression {
		if exp == '+' || exp == '-' || exp == '*' {
			expLeft, expRight := expression[:i], expression[i+1:]
			leftRec, rightRec := diffWaysToCompute(expLeft), diffWaysToCompute(expRight)
			for _, left := range leftRec {
				for _, right := range rightRec {
					switch exp {
					case '+':
						results = append(results, left+right)
					case '-':
						results = append(results, left-right)
					case '*':
						results = append(results, left*right)
					}
				}
			}
		}
	}
	if len(results) == 0 {
		result, _ := strconv.Atoi(expression)
		return []int{result}
	}
	return results
}
