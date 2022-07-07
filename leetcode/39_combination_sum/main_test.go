package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	input    []int
	target   int
	expected [][]int
}{
	{
		input:    []int{2, 3, 6, 7},
		target:   7,
		expected: [][]int{{2, 2, 3}, {7}},
	},
	{
		input:    []int{2, 3, 5},
		target:   8,
		expected: [][]int{{2, 2, 2}, {2, 3, 3}, {3, 5}},
	},
	{
		input:    []int{2},
		target:   1,
		expected: [][]int{},
	},
}

func TestCombinationSum(t *testing.T) {
	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, combinationSum(test.input, test.target))
	}
}
