package main

import (
	"math"
)

func main() {
	// TODO: implement here
}

// Time Limit Exceeded
func findMaxAverageV1(nums []int, k int) float64 {
	if k > len(nums) {
		return 0
	}

	result := float64(math.MinInt)
	for i := 0; i <= len(nums)-k; i++ {
		var total float64
		for _, num := range nums[i : i+k] {
			total += float64(num)
		}
		result = math.Max(result, total/float64(k))
	}
	return result
}

// slide window solution
func findMaxAverageV2(nums []int, k int) float64 {
	if k > len(nums) {
		return 0
	}

	result := float64(math.MinInt)
	var total int
	for _, num := range nums[:k] {
		total += num
	}
	result = math.Max(result, float64(total)/float64(k))
	var prefixIdx int
	for i := k; i < len(nums); i++ {
		total = total - nums[prefixIdx] + nums[i]
		prefixIdx++
		result = math.Max(result, float64(total)/float64(k))
	}
	return result
}

// enhance v2
func findMaxAverageV3(nums []int, k int) float64 {
	if k > len(nums) {
		return 0
	}

	result := float64(math.MinInt)
	var prefix int
	var sum float64
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			sum += float64(nums[i])
		} else {
			sum += float64(nums[i])
			result = math.Max(result, sum/float64(k))
			sum -= float64(nums[prefix])
			prefix++
		}
	}
	return result
}
