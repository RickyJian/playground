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

func productExceptSelfPrefixSum(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	postfix := make([]int, len(nums))
	postfix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}
	results := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		results[i] = prefix[i] * postfix[i]
	}
	return results
}

// TODO: enhance prefix sum
