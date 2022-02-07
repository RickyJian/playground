package main

import (
	"fmt"
)

var nums = []int{-4, -1, 0, 3, 10}
var nums2 = []int{-5, -3, -2, -1}

func main() {
	fmt.Println(sortedSquares(nums))
	fmt.Println(sortedSquares(nums2))
}

func sortedSquares(nums []int) []int {
	numLen := len(nums)
	results := make([]int, numLen)
	l, r := 0, numLen-1
	for i := numLen - 1; i >= 0; i-- {
		left, right := nums[l]*nums[l], nums[r]*nums[r]
		if right >= left {
			results[i] = right
			r--
		} else {
			results[i] = left
			l++
		}
	}
	return results
}
