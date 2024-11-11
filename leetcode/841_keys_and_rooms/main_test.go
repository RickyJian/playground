package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanVisitAllRoomsV1(t *testing.T) {
	var tests = []*struct {
		rooms    [][]int
		expected bool
	}{
		{
			rooms:    [][]int{{1}, {2}, {3}, {}},
			expected: true,
		},
		{
			rooms:    [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canVisitAllRoomsV1(test.rooms), test.expected)
	}
}
