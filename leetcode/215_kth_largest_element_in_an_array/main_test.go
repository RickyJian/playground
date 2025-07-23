package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "範例 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "範例 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "單一元素",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "兩個元素，k=1",
			nums:     []int{2, 1},
			k:        1,
			expected: 2,
		},
		{
			name:     "兩個元素，k=2",
			nums:     []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			name:     "重複元素",
			nums:     []int{3, 3, 3, 3, 3},
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindKthLargestV2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "範例 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "範例 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "單一元素",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "兩個元素，k=1",
			nums:     []int{2, 1},
			k:        1,
			expected: 2,
		},
		{
			name:     "兩個元素，k=2",
			nums:     []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			name:     "重複元素",
			nums:     []int{3, 3, 3, 3, 3},
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargestV2(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindKthLargestV3(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "範例 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "範例 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "單一元素",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "兩個元素，k=1",
			nums:     []int{2, 1},
			k:        1,
			expected: 2,
		},
		{
			name:     "兩個元素，k=2",
			nums:     []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			name:     "重複元素",
			nums:     []int{3, 3, 3, 3, 3},
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargestV3(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
