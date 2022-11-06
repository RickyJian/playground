package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	nums     []int
	left     int
	right    int
	expected int
}{
	{
		nums:     []int{2, 1, 4, 3},
		left:     2,
		right:    3,
		expected: 3,
	},
	{
		nums:     []int{2, 9, 2, 5, 6},
		left:     2,
		right:    8,
		expected: 7,
	},
	{
		nums:     []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52},
		left:     32,
		right:    69,
		expected: 22,
	},
}

func TestNumSubarrayBoundedMaxBruteForce(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, numSubarrayBoundedMaxBruteForce(test.nums, test.left, test.right))
	}
}
