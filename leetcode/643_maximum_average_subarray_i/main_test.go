package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverageV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{-1},
			k:        1,
			expected: -1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, findMaxAverageV1(test.nums, test.k))
	}
}

func TestFindMaxAverageV2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{-1},
			k:        1,
			expected: -1,
		},
		{
			nums:     []int{7, 4, 5, 8, 8, 3, 9, 8, 7, 6},
			k:        7,
			expected: 7,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, findMaxAverageV2(test.nums, test.k), test.nums)
	}
}

func TestFindMaxAverageV3(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{-1},
			k:        1,
			expected: -1,
		},
		{
			nums:     []int{7, 4, 5, 8, 8, 3, 9, 8, 7, 6},
			k:        7,
			expected: 7,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, findMaxAverageV3(test.nums, test.k), test.nums)
	}
}
