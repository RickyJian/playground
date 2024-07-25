package main

func main() {
	// TODO: implement here
}

// Time Limit Exceeded
func longestSubarrayV1(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		subarray := make([]int, len(nums)-1)
		copy(subarray, nums[:i])
		copy(subarray[i:], nums[i+1:])
		var start int
		for j := 0; j < len(subarray); j++ {
			if subarray[j] == 1 {
				if total := j - start + 1; total > max {
					max = total
				}
			} else {
				start = j + 1
			}
		}
	}
	return max
}

func longestSubarrayV2(nums []int) int {
	var start, zeroCount, max int
	for i, num := range nums {
		if num == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		if result := i - start; result > max {
			max = result
		}
	}
	return max
}
