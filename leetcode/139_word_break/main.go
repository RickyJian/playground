package main

func main() {
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	// TODO: add description
	dp[0] = true
	// 將陣列轉成 map 加速比對 substr
	dictSet := make(map[string]struct{})
	for _, word := range wordDict {
		dictSet[word] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// TODO: add description
			if dp[j] {
				if _, ok := dictSet[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	// 當最後一個元素為 true 時，代表字串 s 可以完美被 wordDict 分割
	return dp[len(dp)-1]
}

// TODO: add DFS solution
