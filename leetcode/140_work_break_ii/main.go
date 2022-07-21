package main

import (
	"strings"
)

func main() {
}

func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 {
		return nil
	}

	result := make([]string, 0)
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			words := wordBreak(s[len(word):], wordDict)
			if words == nil {
				result = append(result, word)
			}

			for _, w := range words {
				result = append(result, word+" "+w)
			}
		}
	}
	return result
}

func wordBreakWithMap(s string, wordDict []string) []string {
	return recursion(s, wordDict, make(map[string][]string))
}

func recursion(s string, wordDict []string, cache map[string][]string) []string {
	if len(s) == 0 {
		return nil
	} else if values, ok := cache[s]; ok {
		return values
	}

	result := make([]string, 0)
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			words := recursion(s[len(word):], wordDict, cache)
			if len(words) == 0 {
				result = append(result, word)
			}

			for _, w := range words {
				result = append(result, word+" "+w)
			}
		}
	}
	cache[s] = result
	return result
}
