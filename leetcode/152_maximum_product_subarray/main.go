package main

import (
	"math"
)

func main() {
}

func maxProductBrute(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		if temp > max {
			max = temp
		}
		for j := i + 1; j < len(nums); j++ {
			temp *= nums[j]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func maxProduct(nums []int) int {
	max, maxTemp, minTemp := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp, current := maxTemp, nums[i]
		maxTemp = int(math.Max(float64(current), math.Max(float64(maxTemp*current), float64(minTemp*current))))
		minTemp = int(math.Min(float64(current), math.Min(float64(temp*current), float64(minTemp*current))))
		max = int(math.Max(float64(max), float64(maxTemp)))
	}
	return max
}
