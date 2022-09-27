package main

import (
	"fmt"
)

func main() {
	// fmt.Println(canPartitionDFS([]int{1, 5, 11, 5}))
	fmt.Println(canPartitionDP([]int{1, 5, 11, 5}))
	// fmt.Println(canPartitionDFS([]int{1, 2, 3, 5}))
	// fmt.Println(canPartitionDFS([]int{2, 2, 1, 1}))
	fmt.Println(canPartitionDP([]int{2, 2, 1, 1}))
	// fmt.Println(canPartitionDP([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	// 	100, 100, 100, 100, 100, 100, 99, 97}))
}

func canPartitionDFS(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	if total%2 > 0 {
		return false
	}

	half := total / 2
	for i := 0; i < len(nums); i++ {
		if dfs(nums, i, half, 0) {
			return true
		}
	}
	return false
}

func dfs(nums []int, idx int, target int, current int) bool {
	if target == current {
		return true
	} else if target < current {
		return false
	}

	for i := idx; i < len(nums); i++ {
		if dfs(nums, i+1, target, current+nums[i]) {
			return true
		}
	}
	return false
}

func canPartitionDP(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	if total%2 > 0 {
		return false
	}

	numLen := len(nums)
	target := total / 2
	dp := make([][]bool, numLen)
	for i := 0; i < len(dp); i++ {
		// 因為 target 是從 0 開始因此需要 + 1
		dp[i] = make([]bool, target+1)
	}

	if nums[0] < target {
		// 第一個 元素 < target 代表該值與其餘 nums 有機會加總 = target
		dp[0][nums[0]] = true
	}

	// 由於 num[0] 已設定，直接從 nums[1] ~ nums[numLen-1] 即可
	for i := 1; i < numLen; i++ {
		for j := 0; j <= (target); j++ {
			// TODO: description
			dp[i][j] = dp[i-1][j]
			if nums[i] == j {
				// TODO: description
				dp[i][j] = true
			} else if nums[i] < j {
				// TODO: description
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
			// 已符合條件直接回傳 true
			if dp[i][target] {
				return true
			}
		}
	}
	return dp[numLen-1][target]
}
