package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := math.MinInt64
	var temp int
	for _, num := range nums {
		temp += num
		if max < temp {
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
	}
	return max
}
