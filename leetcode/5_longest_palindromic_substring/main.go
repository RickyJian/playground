package main

func main() {
	// fmt.Println(longestPalindromeBrute("a"))
	// fmt.Println(longestPalindromeBrute("ac"))
	// fmt.Println(longestPalindromeBrute("babad"))
	// fmt.Println(longestPalindromeBrute("cbbd"))
	// fmt.Println(longestPalindromeBrute("aacabdkacaa"))
	// fmt.Println(longestPalindromeDP("babad"))
}

func longestPalindromeBrute(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				subStr := s[i : j+1]
				p := true
				for z := 0; z <= len(subStr)/2; z++ {
					if subStr[z] != subStr[len(subStr)-z-1] {
						p = false
						break
					}
				}
				if p && (len(result) < j+1-i) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}

func longestPalindromeDP(s string) string {
	var result string
	// init array
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1] // 左邊界右移 右邊界左移
				}
			}
			if dp[i][j] && len(result) < len(s[i:j+1]) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func longestPalindromeDP2(s string) string {
	var result string
	// init array
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && len(result) < len(s[i:j+1]) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

// TODO: 中心擴散法
// TODO: Manacher Algorithm
