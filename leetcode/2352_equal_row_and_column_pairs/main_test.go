package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualPairsV1(t *testing.T) {
	var tests = []*struct {
		grid     [][]int
		expected int
	}{
		{
			grid: [][]int{
				{3, 2, 1},
				{1, 7, 6},
				{2, 7, 7},
			},
			expected: 1,
		},
		{
			grid: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
			expected: 3,
		},
		{
			grid: [][]int{
				{13, 13},
				{13, 13},
			},
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, equalPairsV1(test.grid), test.grid)
	}
}
