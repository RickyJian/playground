package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarrayV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			nums:     []int{1, 1, 1},
			expected: 2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, longestSubarrayV1(test.nums), test.nums)
	}
}

func TestLongestSubarrayV2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			nums:     []int{1, 1, 1},
			expected: 2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, longestSubarrayV2(test.nums), test.nums)
	}
}
