package main

func main() {
	// TODO: implement here
}

func pivotIndexV1(nums []int) int {
	leftSums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftSums[i] = 0
		} else {
			leftSums[i] = leftSums[i-1] + nums[i-1]
		}
	}
	rightSums := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightSums[i] = 0
		} else {
			rightSums[i] = rightSums[i+1] + nums[i+1]
		}
	}
	for i, num := range leftSums {
		if rightSums[i] == num {
			return i
		}
	}
	return -1
}

// refer to:
// https://leetcode.com/problems/find-pivot-index/solutions/109255/short-python-o-n-time-o-1-space-with-explanation/?envType=study-plan-v2&envId=leetcode-75
func pivotIndexV2(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	var left int
	for i, num := range nums {
		total -= num
		if left == total {
			return i
		}
		left += num
	}
	return -1
}
