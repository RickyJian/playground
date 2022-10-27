package main

func main() {
}

func countSubstringsN5(s string, t string) int {
	sLen, tLen := len(s), len(t)
	var result int
	for i := 0; i < sLen; i++ {
		for j := i + 1; j <= sLen; j++ {
			for k := 0; k < tLen; k++ {
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
