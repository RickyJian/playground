package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	robot    []int
	factory  [][]int
	expected int
}{
	{
		robot: []int{0, 4, 6},
		factory: [][]int{
			{2, 2},
			{6, 2},
		},
		expected: 4,
	},
	{
		robot: []int{1, -1},
		factory: [][]int{
			{-2, 1},
			{2, 1},
		},
		expected: 2,
	},
}

func TestMinimumTotalDistanceBruteForce(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, minimumTotalDistanceBruteForce(test.robot, test.factory))
	}
}
