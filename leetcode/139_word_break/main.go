package main

import (
	"strings"
)

func main() {
}

func wordBreakDP(s string, wordDict []string) bool {
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

// Time Limit Exceeded
func wordBreakDFS(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s, word) && wordBreakDFS(s[len(word):], wordDict) {
			return true
		}
	}
	return false
}

// Time Limit Exceeded
func wordBreakDFS2(s string, wordDict []string) bool {
	cache := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		cache[word] = struct{}{}
	}
	return recursion(s, cache)
}

func recursion(s string, cache map[string]struct{}) bool {
	if len(s) == 0 {
		return true
	} else if _, ok := cache[s]; ok {
		return true
	}

	for i := range s {
		if _, ok := cache[s[:i]]; ok && recursion(s[i:], cache) {
			return true
		}
	}
	return false
}
