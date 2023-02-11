package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	n        int
	dislikes [][]int
	expected bool
}{
	{
		n:        4,
		dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
		expected: true,
	},
	{
		n:        3,
		dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
		expected: false,
	},
	{
		n:        5,
		dislikes: [][]int{{1, 2}, {3, 4}, {4, 5}, {3, 5}},
		expected: false,
	},
}

func TestPossibleBipartitionDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, possibleBipartitionDFS(test.n, test.dislikes))
	}
}

func TestPossibleBipartitionUF(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, possibleBipartitionUF(test.n, test.dislikes))
	}
}
