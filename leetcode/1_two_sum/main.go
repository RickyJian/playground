package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for i, num := range nums {
		if previous, ok := resultMap[target-num]; ok {
			return []int{previous, i}
		}
		resultMap[num] = i
	}
	return []int{}
}
