package main

func main() {
}

func countSubstringsN5(s string, t string) int {
	sLen, tLen := len(s), len(t)
	var result int
	for i := 0; i < sLen; i++ {
		for j := i + 1; j <= sLen; j++ {
			for k := 0; k < tLen; k++ {
			LOOP:
				for l := k + 1; l <= tLen; l++ {
					// 僅須比較長度一致的 s[i:j], t[l:k]
					if j-i != l-k {
						continue
					}

					var diff int
					for z := 0; z < j-i; z++ {
						// 計算不一致的個數
						// s 由第 i 個位置開始找
						// t 由第 k 個位置開始找
						if s[i+z] != t[k+z] {
							diff++
						}
						if diff > 1 {
							break LOOP
						}
					}
					if diff == 1 {
						result++
					}
				}
			}
		}
	}
	return result
}

func countSubstringsN3(s string, t string) int {
	sLen, tLen := len(s), len(t)
	var result int
	for i := 0; i < sLen; i++ {
		for j := 0; j < tLen; j++ {
			var diff int
			// p: s, t 最後 index 的位置，不可超過 sLen 和 tLen
			for p := 0; i+p < sLen && j+p < tLen; p++ {
				if s[i+p] != t[j+p] {
					diff++
				}
				if diff > 1 {
					break
				} else if diff == 1 {
					result++
				}
			}
		}
	}
	return result
}

func countSubstringsDP(s string, t string) int {
	sLen, tLen := len(s), len(t)
	dp := make([][][]int, sLen)
	for i := range dp {
		dp[i] = make([][]int, tLen)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	var result int
	for i := 0; i < sLen; i++ {
		if s[i] == t[0] {
			dp[i][0][0] = 1
		} else {
			dp[i][0][1] = 1
		}
		result += dp[i][0][1]
	}
	for i := 1; i < tLen; i++ {
		if s[0] == t[i] {
			dp[0][i][0] = 1
		} else {
			dp[0][i][1] = 1
		}
		result += dp[0][i][1]
	}
	for i := 1; i < sLen; i++ {
		for j := 1; j < tLen; j++ {
			if s[i] == t[j] {
				dp[i][j][0] = dp[i-1][j-1][0] + 1
				dp[i][j][1] = dp[i-1][j-1][1]
			} else {
				dp[i][j][0] = 0
				dp[i][j][1] = dp[i-1][j-1][0] + 1
			}
			result += dp[i][j][1]
		}
	}
	return result
}

func countSubstringsDP2(s string, t string) int {
	sLen, tLen := len(s), len(t)
	// dp[i][j][0]：在 s 字串中第 i 個位置與 t 字串中第 j 個位置是否相同
	// dp[i][j][1]：在 s 字串中第 i 個位置與 t 字串中第 j 個位置是否不同
	dp := make([][][]int, sLen+1)
	for i := range dp {
		dp[i] = make([][]int, tLen+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	var result int
	for i := 0; i < sLen; i++ {
		for j := 0; j < tLen; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1][0] = dp[i][j][0] + 1
				dp[i+1][j+1][1] = dp[i][j][1]
			} else {
				dp[i+1][j+1][0] = 0
				//
				dp[i+1][j+1][1] = dp[i][j][0] + 1
			}
			result += dp[i+1][j+1][1]
		}
	}
	return result
}
