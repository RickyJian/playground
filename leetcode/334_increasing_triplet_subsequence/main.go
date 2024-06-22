package main

import (
	"math"
)

func main() {
	// TODO: implement here
}

// result: Time Limit Exceeded
func increasingTripletBrute(nums []int) bool {
	if len(nums) < 3 {
		return false
	} else if len(nums) == 3 {
		return nums[0] < nums[1] && nums[1] < nums[2]
	}
	for i := 0; i < len(nums)-2; i++ {
		left, middle := i, i+1
		// left
		if nums[left] >= nums[middle] {
			continue
		}
		for ; middle < len(nums)-1; middle++ {
			right := middle + 1
			for ; right < len(nums); right++ {
				if nums[middle] >= nums[right] {
					continue
				}
				if nums[left] < nums[middle] && nums[middle] < nums[right] {
					return true
				}
			}
		}
	}
	return false
}

func increasingTripletOn(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}
