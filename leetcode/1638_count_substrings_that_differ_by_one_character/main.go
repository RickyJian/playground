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
					if j-i != l-k {
						continue
					}

					var diff int
					for z := 0; z < j-i; z++ {
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
			for p := 0; i+p < sLen && j+p < tLen; p++ {
				if s[i+p] != t[j+p] {
					diff++
				}
				if diff == 1 {
					result++
				}
				if diff > 1 {
					break
				}
			}
		}
	}
	return result
}
