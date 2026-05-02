package main

import (
	"cmp"
	"slices"
)

func main() {
	// do nothing
}

func maxSubsequence(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	pairs := make([]*pair, len(nums))
	for i, num := range nums {
		pairs[i] = &pair{
			index: i,
			num:   num,
		}
	}
	slices.SortFunc(pairs, func(a, b *pair) int {
		return cmp.Or(
			-cmp.Compare(a.num, b.num),
			cmp.Compare(a.index, b.index),
		)
	})
	pairs = pairs[:k]
	slices.SortFunc(pairs, func(a, b *pair) int {
		return cmp.Compare(a.index, b.index)
	})
	results := make([]int, k)
	for i, pair := range pairs {
		results[i] = pair.num
	}
	return results
}

type pair struct {
	index int
	num   int
}
