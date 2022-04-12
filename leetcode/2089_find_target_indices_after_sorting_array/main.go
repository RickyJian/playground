package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(targetIndicesBruteLoop([]int{1, 2, 5, 2, 3}, 2))
	fmt.Println(targetIndicesBruteLoop([]int{1, 2, 5, 2, 3}, 3))
	fmt.Println(targetIndicesBruteLoop([]int{1, 2, 5, 2, 3}, 5))
	fmt.Println(targetIndicesBruteBS([]int{1, 2, 5, 2, 3}, 2))
	fmt.Println(targetIndicesBruteBS([]int{1, 2, 5, 2, 3}, 3))
	fmt.Println(targetIndicesBruteBS([]int{1, 2, 5, 2, 3}, 5))
}

func targetIndicesBruteLoop(nums []int, target int) []int {
	sort.Ints(nums)
	var result []int
	for i, num := range nums {
		if target == num {
			result = append(result, i)
		}
	}
	return result
}

func targetIndicesBruteBS(nums []int, target int) []int {
	sort.Ints(nums)
	first, last := findFirst(nums, target), findLast(nums, target)
	if first == -1 {
		return []int{}
	}

	result := make([]int, last-first+1)
	for i := 0; i <= last-first; i++ {
		result[i] = first + i
	}
	return result
}

func findFirst(nums []int, target int) int {
	numLen := len(nums)
	left, right := 0, numLen-1
	result := -1
	for left <= right {
		midIdx := (left + right) / 2
		mid := nums[midIdx]
		if target == mid {
			right = midIdx - 1
			result = midIdx
		} else if target < mid {
			right = midIdx - 1
		} else {
			// target > mid
			left = midIdx + 1
		}
	}
	return result
}

func findLast(nums []int, target int) int {
	numLen := len(nums)
	left, right := 0, numLen-1
	result := -1
	for left <= right {
		midIdx := (left + right) / 2
		mid := nums[midIdx]
		if target == mid {
			left = midIdx + 1
			result = midIdx
		} else if target < mid {
			right = midIdx - 1
		} else {
			// target > mid
			left = midIdx + 1
		}
	}
	return result
}
