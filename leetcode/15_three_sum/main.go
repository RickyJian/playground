package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("results:", threeSum([]int{-1, 0, 1}))
	fmt.Println("results:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	numLen := len(nums)
	if numLen <= 1 {
		return [][]int{}
	}

	sort.Ints(nums)

	var results [][]int
	for i := 0; i < numLen; i++ {
		first := nums[i]
		if i > 0 && first == nums[i-1] {
			continue
		}

		left, right := i+1, numLen-1

		for left < right {
			second, last := nums[left], nums[right]
			if sum := first + second + last; sum == 0 {
				results = append(results, []int{first, second, last})
				right--
				for left < right && last == nums[right] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				// sum < 0
				left++
			}
		}
	}
	return results
}
