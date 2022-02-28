package main

import (
	"fmt"
)

func main() {
	// fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	// fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 4, 2, 3}, 1, 1))
	// fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	// fmt.Println(containsNearbyAlmostDuplicate([]int{2147483647, -1, 2147483647}, 1, 2147483647))
	// fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2}, 0, 1))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 0 || k < 0 || t < 0 {
		return false
	}

	bucketWidth := t + 1
	numMap := make(map[int]int)
	for i, num := range nums {
		var bucketIdx int
		if num >= 0 {
			bucketIdx = num / bucketWidth
		} else {
			bucketIdx = num/bucketWidth - 1
		}
		lowerIdx, upperIdx := bucketIdx-1, bucketIdx+1
		if _, ok := numMap[bucketIdx]; ok {
			return true
		} else if lower, ok := numMap[lowerIdx]; ok && abs(num-lower) <= t {
			return true
		} else if upper, ok := numMap[upperIdx]; ok && abs(num-upper) <= t {
			return true
		}
		numMap[bucketIdx] = num
		if i >= k {
			delete(numMap, nums[i-k]/bucketWidth)
		}
	}
	return false
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// TODO: binary search