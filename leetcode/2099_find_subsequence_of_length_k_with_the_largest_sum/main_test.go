package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumOf(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sortedCopy(nums []int) []int {
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Ints(cp)
	return cp
}

func TestMaxSubsequence(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		k           int
		expectedSum int
		expectedLen int
	}{
		{
			name:        "範例 1",
			nums:        []int{2, 1, 3, 3},
			k:           2,
			expectedSum: 6,
			expectedLen: 2,
		},
		{
			name:        "範例 2",
			nums:        []int{-1, -2, 3, 4},
			k:           3,
			expectedSum: 6,
			expectedLen: 3,
		},
		{
			name:        "範例 3",
			nums:        []int{3, 4, 3, 3},
			k:           2,
			expectedSum: 7,
			expectedLen: 2,
		},
		{
			name:        "k 等於陣列長度",
			nums:        []int{1, 2, 3},
			k:           3,
			expectedSum: 6,
			expectedLen: 3,
		},
		{
			name:        "單一元素",
			nums:        []int{5},
			k:           1,
			expectedSum: 5,
			expectedLen: 1,
		},
		{
			name:        "全部為負數",
			nums:        []int{-5, -3, -1, -4},
			k:           2,
			expectedSum: -4,
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubsequence(tt.nums, tt.k)
			assert.Equal(t, tt.expectedLen, len(result), "長度應為 %d", tt.k)
			assert.Equal(t, tt.expectedSum, sumOf(result), "總和應為 %d", tt.expectedSum)

			// 驗證結果中的元素與 nums 前 k 大元素吻合（排序後比較）
			sortedNums := sortedCopy(tt.nums)
			topK := sortedNums[len(sortedNums)-tt.k:]
			assert.Equal(t, sortedCopy(topK), sortedCopy(result), "應選出前 %d 大的元素", tt.k)
		})
	}
}

func TestMaxSubsequenceMinHeap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{
			name: "範例 1",
			nums: []int{2, 1, 3, 3},
			k:    2,
		},
		{
			name: "範例 2",
			nums: []int{-1, -2, 3, 4},
			k:    3,
		},
		{
			name: "範例 3",
			nums: []int{3, 4, 3, 3},
			k:    2,
		},
		{
			name: "k 等於陣列長度",
			nums: []int{1, 2, 3},
			k:    3,
		},
		{
			name: "單一元素",
			nums: []int{5},
			k:    1,
		},
		{
			name: "全部為負數",
			nums: []int{-5, -3, -1, -4},
			k:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := maxSubsequence(tt.nums, tt.k)
			result := maxSubsequenceMinHeap(tt.nums, tt.k)
			assert.Equal(t, expected, result)
		})
	}
}
