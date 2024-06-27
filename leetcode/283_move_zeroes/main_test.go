package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZerosV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
	}
	for _, test := range tests {
		moveZeroesV1(test.nums)
		assert.Equal(t, test.expected, test.nums)
	}
}

func TestMoveZerosV2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
	}
	for _, test := range tests {
		moveZeroesV2(test.nums)
		assert.Equal(t, test.expected, test.nums)
	}
}
