package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	isConnected [][]int
	expected    int
}{
	{
		isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		expected:    2,
	},
	{
		isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		expected:    3,
	},
	{
		isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
		expected:    1,
	},
	{
		isConnected: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		expected:    1,
	},
}

func TestFindCircleNum(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, findCircleNum(test.isConnected))
	}
}

func TestFindCircleNumV2(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, findCircleNumV2(test.isConnected))
	}
}
