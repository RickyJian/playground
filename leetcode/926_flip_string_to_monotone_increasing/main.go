package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlipsMonoIncrBrute("00110"))
	fmt.Println(minFlipsMonoIncrBrute("010110"))
	fmt.Println(minFlipsMonoIncrBrute("00011000"))
}

// TLE
func minFlipsMonoIncrBrute(s string) int {
	min := -1
	length := len(s)
	for i := 0; i <= length; i++ {
		var count int
		// right
		for j := i; j < length; j++ {
			if s[j] == '0' {
				count++
			}
		}
		// left
		for j := i - 1; i > 0 && j >= 0; j-- {
			if s[j] == '1' {
				count++
			}
		}
		if min == -1 || min > count {
			min = count
		}
	}
	return min
}
