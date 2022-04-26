package main

import (
	"math"
	"sort"
)

func main() {
}

func smallestDistancePairBrute(nums []int, k int) int {
	results := make([]int, 0, 1000000)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			results = append(results, int(math.Abs(float64(nums[j]-nums[i]))))
		}
	}
	sort.Ints(results)
	return results[k-1]
}

func smallestDistancePairBrute2(nums []int, k int) int {
	counts := make([]int, 1000000)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result := int(math.Abs(float64(nums[j] - nums[i])))
			counts[result]++
		}
	}
	var temp int
	for idx, count := range counts {
		temp += count
		if temp >= k {
			return idx
		}
	}
	return -1
}

func smallestDistancePairBS(nums []int, k int) int {
	sort.Ints(nums)

	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)/2
		var count, left int
		for right, num := range nums {
			for num-nums[left] > mid {
				left++
			}
			count += right - left
		}
		if count < k {
			low = mid + 1
		} else {
			// count >= k
			high = mid
		}
	}
	return low
}
