package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProductBrute(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2},
			expected: -2,
		},
		{
			nums:     []int{0, 2},
			expected: 2,
		},
		{
			nums:     []int{-4, 0, -1, -2, -5, -7},
			expected: 70,
		},
		{
			nums:     []int{7, -2, -4},
			expected: 56,
		},
		{
			nums:     []int{-3, 0, 1, -2},
			expected: 1,
		},
		{
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			nums:     []int{0},
			expected: 0,
		},
	}
	for _, test := range tests {
		result := maxProductBrute(test.nums)
		assert.Equal(t, test.expected, result)
	}
}

func TestMaxProduct(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2},
			expected: -2,
		},
		{
			nums:     []int{0, 2},
			expected: 2,
		},
		{
			nums:     []int{-4, 0, -1, -2, -5, -7},
			expected: 70,
		},
		{
			nums:     []int{7, -2, -4},
			expected: 56,
		},
		{
			nums:     []int{-3, 0, 1, -2},
			expected: 1,
		},
		{
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			nums:     []int{0},
			expected: 0,
		},
	}
	for _, test := range tests {
		result := maxProduct(test.nums)
		assert.Equal(t, test.expected, result)
	}
}
