package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected [][]int
	}{
		{
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{-1, 0, 1},
			},
		},
		{
			nums: []int{0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			nums:     []int{},
			expected: [][]int{},
		},
		{
			nums:     []int{0},
			expected: [][]int{},
		},
		{
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			expected: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, test := range tests {
		results := threeSum(test.nums)
		assert.Equal(t, test.expected, results)
	}
}
