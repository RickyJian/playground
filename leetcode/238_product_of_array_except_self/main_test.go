package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelfBrute(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for _, test := range tests {
		assert.Exactly(t, test.expected, productExceptSelfBrute(test.nums))
	}
}

func TestProductExceptPrefixSum(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for _, test := range tests {
		assert.Exactly(t, test.expected, productExceptSelfPrefixSum(test.nums))
	}
}
