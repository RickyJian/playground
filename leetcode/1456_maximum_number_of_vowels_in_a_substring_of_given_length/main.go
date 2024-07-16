package main

func main() {
	// TODO: implement here
}

// Time Limit Exceeded
func maxVowelsV1(s string, k int) int {
	vowelSet := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	var result int
	for i := 0; i <= len(s)-k; i++ {
		var count int
		for _, word := range []byte(s[i : i+k]) {
			if _, ok := vowelSet[word]; ok {
				count++
			}
			if count == k {
				return count
			}
		}
		if count > result {
			result = count
		}
	}
	return result
}

func maxVowelsV2(s string, k int) int {
	vowelSet := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	var result, sum, left int
	for i := 0; i < len(s); i++ {
		if _, ok := vowelSet[s[i]]; ok {
			sum++
		}
		if i > k-1 {
			if _, ok := vowelSet[s[left]]; ok {
				sum--
			}
			left++
		}
		if result == k {
			return k
		} else if result < sum {
			result = sum
		}
	}
	return result
}

// use function `isVowel` instead of map
func maxVowelsV3(s string, k int) int {
	var result, sum, left int
	for i := 0; i < len([]byte(s)); i++ {
		if isVowel(s[i]) {
			sum++
		}
		if i > k-1 {
			if isVowel(s[left]) {
				sum--
			}
			left++
		}
		if result == k {
			return k
		} else if result < sum {
			result = sum
		}
	}
	return result
}

func isVowel(word byte) bool {
	return word == 'a' || word == 'e' || word == 'i' || word == 'o' || word == 'u'
}
