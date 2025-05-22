package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinReorder(t *testing.T) {
	var tests = []*struct {
		n           int
		connections [][]int
		expected    int
	}{
		{
			n:           6,
			connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			expected:    3,
		},
		{
			n:           5,
			connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			expected:    2,
		},
		{
			n:           3,
			connections: [][]int{{1, 0}, {2, 0}},
			expected:    0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, minReorder(test.n, test.connections))
	}
}
