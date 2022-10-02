package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartitionDFS(t *testing.T) {
	var tests = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{2, 2, 1, 1},
			expected: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canPartitionDFS(test.nums))
	}
}

func TestCanPartitionDFSMemo(t *testing.T) {
	var tests = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{2, 2, 1, 1},
			expected: true,
		},
		{
			nums: []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 99, 97},
			expected: false,
		},
		{
			nums:     []int{1, 13, 19, 17, 17, 15, 14, 12, 4},
			expected: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canPartitionDFSMemo(test.nums))
	}
}

func TestCanPartitionDP(t *testing.T) {
	var tests = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{2, 2, 1, 1},
			expected: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canPartitionDP(test.nums))
	}
}

func TestCanPartitionDP2(t *testing.T) {
	var tests = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{2, 2, 1, 1},
			expected: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canPartitionDP2(test.nums))
	}
}
