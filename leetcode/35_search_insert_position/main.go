package main

import (
	"fmt"
)

var nums = []int{1, 3, 5, 7}

func main() {
	fmt.Println(searchInsert(nums, 0))
	fmt.Println(searchInsert(nums, 1))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 3))
	fmt.Println(searchInsert(nums, 4))
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 6))
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 8))
}

func searchInsert(nums []int, target int) int {
	numLen := len(nums)
	if len(nums) == 0 || target < nums[0] {
		return 0
	} else if target > nums[numLen-1] {
		return numLen
	}

	start, end := 0, numLen-1
	for {
		mid := (start + end) / 2
		midNum := nums[mid]
		if target == midNum || len(nums[start:end]) == 0 {
			return mid
		} else if target > midNum {
			start = mid + 1
		} else {
			end = mid
		}
	}
}
