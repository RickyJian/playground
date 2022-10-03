package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numRollsToTargetDP(1, 6, 3))
	fmt.Println(numRollsToTargetDP(2, 6, 7))
}

func numRollsToTargetDP(n int, k int, target int) int {
	m := int(math.Pow10(9) + 7)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for z := j; z <= target; z++ {
				dp[i][z] = (dp[i][z] + dp[i-1][z-j]) % m
			}
		}
	}
	return dp[n][target]
}
