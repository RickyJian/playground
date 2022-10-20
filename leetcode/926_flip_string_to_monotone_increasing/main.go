package main

import (
	"fmt"
)

func main() {
	// fmt.Println(minFlipsMonoIncrBrute("00110"))
	// fmt.Println(minFlipsMonoIncrBrute("010110"))
	// fmt.Println(minFlipsMonoIncrBrute("00011000"))
	fmt.Println(minFlipsMonoIncrBrute("11011"))

	fmt.Println(minFlipsMonoIncrDP1("00110"))
	// fmt.Println(minFlipsMonoIncrDP1("010110"))
	// fmt.Println(minFlipsMonoIncrDP1("00011000"))
	// fmt.Println(minFlipsMonoIncrDP1("11011"))
}

// TLE
func minFlipsMonoIncrBrute(s string) int {
	min := -1
	length := len(s)
	for i := 0; i <= length; i++ {
		var count int
		// flip to one
		for j := i; j < length; j++ {
			if s[j] == '0' {
				count++
			}
		}
		// flip to zero
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

func minFlipsMonoIncrDP1(s string) int {
	length := len(s)
	flipToOne := make(map[int]int, length)
	for i := length - 1; i >= 0; i-- {
		flipToOne[i] = flipToOne[i+1]
		if s[i] == '0' {
			flipToOne[i]++
		}
	}
	flipToZero := make(map[int]int, length)
	for i := 0; i < length; i++ {
		flipToZero[i] = flipToZero[i-1]
		if s[i] == '1' {
			flipToZero[i]++
		}
	}
	min := -1
	for i := 0; i < length; i++ {
		var count int
		if i == 0 {
			// 左邊第一數右邊皆須翻成 1
			count = flipToOne[i]
		} else if i == length-1 {
			// 右邊第一數右邊皆須翻成 0
			count = flipToZero[i]
		} else {
			// TODO: description
			count = flipToZero[i] + flipToOne[i+1]
		}
		if min == -1 || min > count {
			min = count
		}
	}
	return min
}

// TODO: enhance dp
