package main

import (
	"fmt"
)

func main() {
	// Brute force solution
	// fmt.Println(minFlipsMonoIncrBrute("00110"))
	// fmt.Println(minFlipsMonoIncrBrute("010110"))
	// fmt.Println(minFlipsMonoIncrBrute("00011000"))
	fmt.Println(minFlipsMonoIncrBrute("11011"))

	// DP two arrays solution
	fmt.Println(minFlipsMonoIncrDP2Array("00110"))
	// fmt.Println(minFlipsMonoIncrDP2Array("010110"))
	// fmt.Println(minFlipsMonoIncrDP2Array("00011000"))
	// fmt.Println(minFlipsMonoIncrDP2Array("11011"))

	// DP one array solution
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

func minFlipsMonoIncrDP2Array(s string) int {
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
	var min int
	for i := 0; i <= length; i++ {
		// 當執行到第 i 個元素時：
		// flipToZero[i-1]: 左邊(i-1)需要翻轉成 0 的次數
		// flipToOne[i]: 右邊(i)需要翻轉成 1 的次數
		count := flipToZero[i-1] + flipToOne[i]
		if i == 0 || min > count {
			min = count
		}
	}
	return min
}

func minFlipsMonoIncrDP1Array(s string) int {
	length := len(s)
	flipToZero := make(map[int]int)
	for i := 0; i < length; i++ {
		flipToZero[i] = flipToZero[i-1]
		if s[i] == '1' {
			flipToZero[i]++
		}
	}

	var min int
	for i := 0; i <= length; i++ {
		// flipToZero[i-1]:
		// right: ((length - i) - (flipToZero[length-1] - flipToZero[i-1]))
		//   * length - i: TODO: add description
		//   * flipToZero[length-1]: 字串最後一個元素
		//   * flipToZero[length-1] - flipToZero[i-1]: 從 i ~ 最後元素須翻轉的次數
		count := flipToZero[i-1] + ((length - i) - (flipToZero[length-1] - flipToZero[i-1]))
		if i == 0 || count < min {
			min = count
		}
	}
	return min
}
