package main

import (
	"fmt"
)

var (
	input   = []int{-1, 0, 3, 5, 9, 12}
	target1 = 9
	target2 = 2
)

func main() {
	// Output: 4
	fmt.Println(search(input, target1))
	// Output: -1
	fmt.Println(search(input, target2))
}

func search(nums []int, target int) int {
	index := len(nums) / 2
	value := nums[index]
	if len(nums) == 0 ||
		target < nums[0] ||
		target > nums[len(nums)-1] {
		return -1
	} else if target > value {
		result := search(nums[index+1:], target)
		// 若查無資料直接回傳 -1
		if result == -1 {
			return result
		}
		// 查有資料須加上原本的中間值
		return index + result + 1
	} else if target < value {
		return search(nums[:index], target)
	} else {
		return index
	}
}
