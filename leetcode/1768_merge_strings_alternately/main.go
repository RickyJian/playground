package main

func main() {
	// TODO: implement here
}

func mergeAlternately(word1 string, word2 string) string {
	r1, r2 := []rune(word1), []rune(word2)
	w1Len, w2Len := len(r1), len(r2)
	size := w1Len
	if size < w2Len {
		size = w2Len
	}
	result := make([]rune, 0, w1Len+w2Len)
	for i := 0; i < size; i++ {
		if w1Len > i {
			result = append(result, r1[i])
		}
		if w2Len > i {
			result = append(result, r2[i])
		}
	}
	return string(result)
}
