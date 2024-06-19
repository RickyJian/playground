package main

func main() {
	// TODO: implement here
}

// result: Time Limit Exceeded
func productExceptSelfBrute(nums []int) []int {
	numLen := len(nums)
	results := make([]int, numLen)
	for i := 0; i < numLen; i++ {
		results[i] = 1
		for j := 0; j < numLen; j++ {
			if i == j {
				continue
			}
			results[i] *= nums[j]
		}
	}
	return results
}
