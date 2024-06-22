package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTripletBrute(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			nums:     []int{20, 100, 10, 12, 5, 13},
			expected: true,
		},
		{
			nums:     []int{0, 4, 2, 1, 0, -1, -3},
			expected: false,
		},
		{
			nums:     []int{1, 2, 2, 1},
			expected: false,
		},
		{
			nums:     []int{1, 5, 0, 4, 1, 3},
			expected: true,
		},
		{
			nums:     []int{6, 7, 1, 2},
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, increasingTripletBrute(test.nums), test.nums)
	}
}

func TestIncreasingTripletOn(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			nums:     []int{20, 100, 10, 12, 5, 13},
			expected: true,
		},
		{
			nums:     []int{0, 4, 2, 1, 0, -1, -3},
			expected: false,
		},
		{
			nums:     []int{1, 2, 2, 1},
			expected: false,
		},
		{
			nums:     []int{1, 5, 0, 4, 1, 3},
			expected: true,
		},
		{
			nums:     []int{6, 7, 1, 2},
			expected: false,
		},
		{
			nums:     []int{10, 12, 5, 1, 2},
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, increasingTripletOn(test.nums), test.nums)
	}
}
