package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{-1, 100, 2, 100, 100, 4, 100},
			expected: 100,
		},
	}
	for _, test := range tests {
		results := majorityElement(test.nums)
		assert.Equal(t, test.expected, results)
	}
}

func TestMajorityElement2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{-1, 100, 2, 100, 100, 4, 100},
			expected: 100,
		},
	}
	for _, test := range tests {
		results := majorityElement2(test.nums)
		assert.Equal(t, test.expected, results)
	}
}
