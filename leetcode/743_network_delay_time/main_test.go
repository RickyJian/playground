package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	times    [][]int
	n        int
	k        int
	expected int
}{
	{
		times: [][]int{
			{2, 1, 1},
			{2, 3, 1},
			{3, 4, 1},
		},
		n:        4,
		k:        2,
		expected: 2,
	},
	{
		times: [][]int{
			{1, 2, 1},
		},
		n:        2,
		k:        1,
		expected: 1,
	},
	{
		times: [][]int{
			{1, 2, 1},
		},
		n:        2,
		k:        2,
		expected: -1,
	},
	{
		times: [][]int{
			{1, 2, 1},
			{2, 1, 3},
		},
		n:        2,
		k:        2,
		expected: 3,
	},
}

func TestNetworkDelayTimeDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, networkDelayTime(test.times, test.n, test.k))
	}
}
