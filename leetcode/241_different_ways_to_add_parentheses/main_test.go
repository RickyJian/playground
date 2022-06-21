package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	expressions string
	expected    []int
}{
	{
		expressions: "2-1-1",
		expected:    []int{0, 2},
	},
	{
		expressions: "2*3-4*5",
		expected:    []int{-34, -14, -10, -10, 10},
	},
	{
		expressions: "2",
		expected:    []int{2},
	},
}

func TestDiffWaysToCompute(t *testing.T) {
	for _, test := range tests {
		result := diffWaysToCompute(test.expressions)
		assert.ElementsMatch(t, test.expected, result)
	}
}
