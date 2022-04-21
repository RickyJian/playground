package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	arr      []int
	k        int
	x        int
	expected []int
}{
	{
		arr:      []int{1, 2, 3, 4, 5},
		k:        4,
		x:        3,
		expected: []int{1, 2, 3, 4},
	},
	{
		arr:      []int{1, 2, 3, 4, 5},
		k:        4,
		x:        -1,
		expected: []int{1, 2, 3, 4},
	},
	{
		arr:      []int{1, 1, 1, 10, 10, 10},
		k:        1,
		x:        9,
		expected: []int{10},
	},
}

func TestFindClosestElementsS2E(t *testing.T) {
	for _, test := range tests {
		results := findClosestElementsS2E(test.arr, test.k, test.x)
		assert.Equal(t, test.expected, results)
	}
}

func TestFindClosestElementsBS(t *testing.T) {
	for _, test := range tests {
		results := findClosestElementsBS(test.arr, test.k, test.x)
		assert.Equal(t, test.expected, results)
	}
}
