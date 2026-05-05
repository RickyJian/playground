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

func maxSubsequenceMinHeap(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	minHeap := newMinHeap(k)
	for i, num := range nums {
		minHeap.push(&pair{
			index: i,
			num:   num,
		})
	}
	slices.SortFunc(minHeap.pairs, func(a, b *pair) int {
		return cmp.Compare(a.index, b.index)
	})
	results := make([]int, k)
	for i, pair := range minHeap.pairs {
		results[i] = pair.num
	}
	return results
}

type pair struct {
	index int
	num   int
}

type minHeap struct {
	k     int
	pairs []*pair
}

func newMinHeap(k int) *minHeap {
	return &minHeap{k: k, pairs: make([]*pair, 0, k)}
}

func (h *minHeap) push(p *pair) {
	if len(h.pairs) == h.k {
		// bubble down：當前節點與左右子節點相比，大於做置換
		if top := h.pairs[0]; top.num < p.num || (top.num == p.num && top.index > p.index) {
			h.pairs[0] = p
			var currentIdx int
			for {
				currentNode := h.pairs[currentIdx]
				smallest, smallestIdx := h.pairs[currentIdx], currentIdx
				if leftIdx := 2*currentIdx + 1; leftIdx < len(h.pairs) {
					if leftNode := h.pairs[leftIdx]; leftNode.num < smallest.num ||
						(leftNode.num == smallest.num && leftNode.index > smallest.index) {
						smallest = leftNode
						smallestIdx = leftIdx
					}
				}
				if rightIdx := 2*currentIdx + 2; rightIdx < len(h.pairs) {
					if rightNode := h.pairs[rightIdx]; rightNode.num < smallest.num ||
						(rightNode.num == smallest.num && rightNode.index > smallest.index) {
						smallest = rightNode
						smallestIdx = rightIdx
					}
				}
				if smallest.num == currentNode.num && smallest.index == currentNode.index {
					return
				}
				h.pairs[currentIdx], h.pairs[smallestIdx] = smallest, currentNode
				currentIdx = smallestIdx
			}
		}
	} else {
		// bubble up：當前節點與父節點相比，小於做置換
		h.pairs = append(h.pairs, p)
		currentIdx := len(h.pairs) - 1
		for currentIdx > 0 {
			parentIdx := (currentIdx - 1) / 2
			parentNode := h.pairs[parentIdx]
			currentNode := h.pairs[currentIdx]
			if parentNode.num > currentNode.num ||
				(parentNode.num == currentNode.num && parentNode.index > currentNode.index) {
				h.pairs[parentIdx], h.pairs[currentIdx] = h.pairs[currentIdx], h.pairs[parentIdx]
				currentIdx = parentIdx
			} else {
				return
			}
		}
	}
}
