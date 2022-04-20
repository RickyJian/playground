package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums     []int
	target   int
	expected int
}{
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   4,
		expected: 0,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   5,
		expected: 1,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   6,
		expected: 2,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   7,
		expected: 3,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   1,
		expected: 5,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   2,
		expected: 6,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
	},
}

func TestSearchBrute(t *testing.T) {
	for _, test := range tests {
		result := searchBrute(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestSearchBSPvMid(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   2,
			expected: 0,
		},
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   3,
			expected: 1,
		},
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   4,
			expected: 2,
		},
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   0,
			expected: 3,
		},
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   1,
			expected: 4,
		},
		{
			nums:     []int{2, 3, 4, 0, 1},
			target:   5,
			expected: -1,
		},
	}
	for _, test := range tests {
		result := searchBS(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestSearchBSPvLeft(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   3,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   4,
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   0,
			expected: 4,
		},
		{
			nums:     []int{1, 2, 3, 4, 0},
			target:   5,
			expected: -1,
		},
	}
	for _, test := range tests {
		result := searchBS(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestSearchBSPvRight(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   3,
			expected: 0,
		},
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   4,
			expected: 1,
		},
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   0,
			expected: 2,
		},
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   1,
			expected: 3,
		},
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   2,
			expected: 4,
		},
		{
			nums:     []int{3, 4, 0, 1, 2},
			target:   5,
			expected: -1,
		},
	}
	for _, test := range tests {
		result := searchBS(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}
