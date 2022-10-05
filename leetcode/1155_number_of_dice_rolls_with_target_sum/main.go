package main

import (
	"fmt"
)

func main() {
	fmt.Println(numRollsToTargetDP(1, 6, 3))
	fmt.Println(numRollsToTargetDP(2, 6, 7))
	fmt.Println(numRollsToTargetDP(3, 6, 7))
	fmt.Println("====")
	fmt.Println(numRollsToTargetDFS(1, 6, 3))
	fmt.Println(numRollsToTargetDFS(2, 6, 7))
	fmt.Println(numRollsToTargetDFS(3, 6, 7))
}

func numRollsToTargetDP(n int, k int, target int) int {
	// 若 n 個骰子的最大值都無超過 target 代表骰子終究無法組合成 target
	if n*k < target {
		return 0
	}

	m := 1000000007
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	// dp[i][j]：表示有 i 顆骰子可以組成 j 的方法有幾種
	// 雖然題目不會有 dp[0][0] 的情形發生，但為了讓後續 dp 可以向下執行需要個初始值，
	// 而我們可以看作當有 0 顆骰子可以組成 0 的方法有 1 種
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// 我們只須檢查 z ~ target 有幾種組合方式，
			// 因為 1 ~ (z - 1) 會是我們 j -1 執行的結果
			for z := j; z <= target; z++ {
				// TODO: add description
				// dp[i-1][z-j]：檢查上一顆骰子骰出的結果有幾種方法可以組成 z-j
				dp[i][z] = (dp[i][z] + dp[i-1][z-j]) % m
			}
		}
	}
	return dp[n][target]
}

// TODO: dp compress

func numRollsToTargetDFS(n int, k int, target int) int {
	// 若 n 個骰子的最大值都無超過 target 代表骰子終究無法組合成 target
	if n*k < target {
		return 0
	}
	return dfs(n, k, target)
}

func dfs(n, k, target int) int {
	if target == 0 && n == 0 {
		// 為最後一個骰子且剛好可以組合成 target
		return 1
	} else if target < 0 || n == 0 {
		// 為最後一個骰子 或 target 無法組合
		return 0
	}

	var count int
	for i := 1; i <= k; i++ {
		target -= i
		// n-1：骰子骰過不可再用
		count += dfs(n-1, k, target) % 1000000007
		// 把剛剛 減 的 i 加回，為了不影響後續 backtracking 使用
		target += i
	}
	return count
}

// TODO: dfs with memo
