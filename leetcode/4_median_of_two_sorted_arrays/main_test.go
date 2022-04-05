package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	num1     []int
	num2     []int
	expected float64
}{
	{
		num1:     []int{1, 3},
		num2:     []int{2},
		expected: 2,
	},
	{
		num1:     []int{1, 2},
		num2:     []int{3, 4},
		expected: 2.5,
	},
	{
		num1:     []int{1, 3},
		num2:     []int{2, 4, 6},
		expected: 3.0,
	},
	{
		num1:     []int{1, 3, 5},
		num2:     []int{2, 4, 6},
		expected: 3.5,
	},
	{
		num1:     []int{1, 2, 3},
		num2:     []int{4, 5, 6},
		expected: 3.5,
	},
	{
		num1:     []int{4, 5, 6},
		num2:     []int{1, 2, 3},
		expected: 3.5,
	},
}

func TestFindMedianSortedArraysBrute(t *testing.T) {
	for _, test := range tests {
		result := findMedianSortedArraysBrute(test.num1, test.num2)
		assert.Equal(t, test.expected, result)
	}
}

func TestFindMedianSortedArraysLogMN(t *testing.T) {
	for _, test := range tests {
		result := findMedianSortedArraysLogMN(test.num1, test.num2)
		assert.Equal(t, test.expected, result)
	}
}
