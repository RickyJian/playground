package main

func main() {
	// TODO: implement here
}

func reverseVowels(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; {
		prefix, suffix := bs[i], bs[j]
		preHasVowel, sufHasVowel := isVowels(prefix), isVowels(suffix)
		if preHasVowel && sufHasVowel {
			bs[i], bs[j] = bs[j], bs[i]
			i++
			j--
		}
		if !preHasVowel {
			i++
		}
		if !sufHasVowel {
			j--
		}
	}
	return string(bs)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
