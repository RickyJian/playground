package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesCcanPartitionDFS(t *testing.T) {
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

func TesCcanPartitionDP(t *testing.T) {
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

func TesCcanPartitionDP2(t *testing.T) {
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
