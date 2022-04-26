package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	nums     []int
	k        int
	expected int
}{
	{
		nums:     []int{1, 3, 1},
		k:        1,
		expected: 0,
	},
	{
		nums:     []int{1, 1, 1},
		k:        2,
		expected: 0,
	},
	{
		nums:     []int{1, 6, 1},
		k:        3,
		expected: 5,
	},
	{
		nums:     []int{62, 100, 4},
		k:        2,
		expected: 58,
	},
}

func TestSmallestDistancePairBrute(t *testing.T) {
	for _, test := range tests {
		result := smallestDistancePairBrute(test.nums, test.k)
		assert.Equal(t, test.expected, result)
	}
}
