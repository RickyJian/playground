package main

import "sort"

func main() {
	// TODO: implement here
}

func closeStringsV1(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	const alphabetCount = 26
	arr1 := make([]int, alphabetCount)
	for _, word := range word1 {
		arr1[word-'a']++
	}
	arr2 := make([]int, alphabetCount)
	for _, word := range word2 {
		arr2[word-'a']++
	}

	for i := 0; i < alphabetCount; i++ {
		if (arr1[i] > 0) != (arr2[i] > 0) {
			return false
		}
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] > arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] > arr2[j]
	})
	for i := 0; i < alphabetCount; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
