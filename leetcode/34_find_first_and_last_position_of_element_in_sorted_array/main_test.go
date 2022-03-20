package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRangeBrute(t *testing.T) {
	var tests = []*struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
	}
	for _, test := range tests {
		result := searchRangeBrute(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestSearchRangeBinarySearch(t *testing.T) {
	var tests = []*struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			nums:     []int{5, 7, 8, 8, 8, 10},
			target:   8,
			expected: []int{2, 4},
		},
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
	}
	for _, test := range tests {
		result := searchRangeBinarySearch(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}
