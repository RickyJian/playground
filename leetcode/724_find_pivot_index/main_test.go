package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndexV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			nums:     []int{2, 1, -1},
			expected: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, pivotIndexV1(test.nums), test.nums)
	}
}

func TestPivotIndexV2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			nums:     []int{2, 1, -1},
			expected: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, pivotIndexV2(test.nums), test.nums)
	}
}
