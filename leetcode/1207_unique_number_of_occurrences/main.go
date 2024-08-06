package main

import (
	"sort"
)

func main() {
	// TODO: implement here
}

func uniqueOccurrencesV1(arr []int) bool {
	sort.Ints(arr)
	numSet := make(map[int]struct{})
	var prev int
	for i, num := range arr {
		if i > 0 && num != arr[i-1] {
			count := i - prev
			if _, ok := numSet[count]; ok {
				return false
			}
			prev = i
			numSet[count] = struct{}{}
		}
	}
	if _, ok := numSet[len(arr)-prev]; ok {
		return false
	}
	return true
}

func uniqueOccurrencesV2(arr []int) bool {
	numCounts := make(map[int]int)
	for _, num := range arr {
		numCounts[num]++
	}
	detects := make(map[int]struct{})
	for _, count := range numCounts {
		if _, ok := detects[count]; ok {
			return false
		}
		detects[count] = struct{}{}
	}
	return true
}
