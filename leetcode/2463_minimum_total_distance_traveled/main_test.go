package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	robot    []int
	factory  [][]int
	expected int64
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
	{
		robot: []int{9, 11, 99, 101},
		factory: [][]int{
			{10, 1},
			{7, 1},
			{14, 1},
			{100, 1},
			{96, 1},
			{103, 1},
		},
		expected: 6,
	},
}

func TestMinimumTotalDistanceBruteForce(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, minimumTotalDistanceDFS(test.robot, test.factory))
	}
}
