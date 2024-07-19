package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnesV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 0, 0},
			k:        0,
			expected: 3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, longestOnesV1(test.nums, test.k), test.nums)
	}
}
