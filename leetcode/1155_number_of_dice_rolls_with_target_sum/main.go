package main

import (
	"fmt"
)

func main() {
	// fmt.Println(numRollsToTargetDP(1, 6, 3))
	fmt.Println(numRollsToTargetDP(2, 6, 7))
	// fmt.Println(numRollsToTargetDP(3, 6, 7))
	// fmt.Println(numRollsToTargetDP(30, 30, 500))
	// fmt.Println(numRollsToTargetDP(10, 10, 30))
	fmt.Println("====")
	// fmt.Println(numRollsToTargetDFS(1, 6, 3))
	// fmt.Println(numRollsToTargetDFS(2, 6, 7))
	// fmt.Println(numRollsToTargetDFS(3, 6, 7))
	// fmt.Println(numRollsToTargetDP(30, 30, 500))
	fmt.Println("====")
	// fmt.Println(numRollsToTargetDFSMemo(1, 6, 3))
	// fmt.Println(numRollsToTargetDFSMemo(2, 6, 7))
	// fmt.Println(numRollsToTargetDFSMemo(3, 6, 7))
	// fmt.Println(numRollsToTargetDFSMemo(30, 30, 500))
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
	// i：第 i 顆骰子
	// j：骰出 j 的面
	// 下方三個 for 可以看成：第 i 個骰子骰出 j 面時，有幾種方式可以組成 z
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// 由於骰出的面(j) > z 沒有方法能組成 z，
			// 因此我們只須檢查疼組成 (z ~ target) 的方法就好
			for z := j; z <= target; z++ {
				// dp[i-1][z-j]：檢查上一顆骰子是否能組成 z-j，若能代表當下的骰子也可以組成 z
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
		// n-1：骰子骰過不可再用
		count += dfs(n-1, k, target-i) % 1000000007
	}
	return count
}

func numRollsToTargetDFSMemo(n int, k int, target int) int {
	// 若 n 個骰子的最大值都無超過 target 代表骰子終究無法組合成 target
	if n*k < target {
		return 0
	}
	return dfsMemo(n, k, target, make(map[string]int))
}

func dfsMemo(n, k, target int, memo map[string]int) int {
	key := fmt.Sprintf("%d_%d", n, target)
	if target == 0 && n == 0 {
		// 為最後一個骰子且剛好可以組合成 target
		return 1
	} else if target < 0 || n == 0 {
		// 為最後一個骰子 或 target 無法組合
		return 0
	} else if val, ok := memo[key]; ok {
		// memo[key] 保存 dfs 曾經走過路的結果，
		// 若有紀錄過可以直接能 memo 中的值，
		// 來減少 dfs 重複走過路徑的情形
		return val
	}

	for i := 1; i <= k; i++ {
		// n-1：骰子骰過不可再用
		memo[key] = (memo[key] + dfsMemo(n-1, k, target-i, memo)) % 1000000007
	}
	return memo[key]
}
