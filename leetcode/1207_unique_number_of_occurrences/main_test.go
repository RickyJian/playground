package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrencesV1(t *testing.T) {
	var tests = []*struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2},
			expected: false,
		},
		{
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
		{
			arr:      []int{-4, 3, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2, 3, 3},
			expected: false,
		},
	}
	for _, test := range tests {
		nums := make([]int, len(test.arr))
		copy(nums, test.arr)
		assert.Equal(t, test.expected, uniqueOccurrencesV1(test.arr), nums)
	}
}

func TestUniqueOccurrencesV2(t *testing.T) {
	var tests = []*struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2},
			expected: false,
		},
		{
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
		{
			arr:      []int{-4, 3, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2, 3, 3},
			expected: false,
		},
	}
	for _, test := range tests {
		nums := make([]int, len(test.arr))
		copy(nums, test.arr)
		assert.Equal(t, test.expected, uniqueOccurrencesV2(test.arr), nums)
	}
}
