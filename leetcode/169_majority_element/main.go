package main

import (
	"fmt"
)

func main() {
	fmt.Println("results:", majorityElement([]int{3, 2, 3}))
	fmt.Println("results:", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println("results:", majorityElement([]int{-1, 100, 2, 100, 100, 4, 100}))
}

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
		if numMap[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

// refer to: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElement2(nums []int) int {
	var n, count int
	for _, num := range nums {
		if count == 0 {
			n, count = num, 1
		} else if n == num {
			count++
		} else {
			count--
		}
	}
	return n
}
