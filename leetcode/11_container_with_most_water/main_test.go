package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaV1(t *testing.T) {
	var tests = []*struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxAreaV1(test.height))
	}
}

func TestMaxAreaV2(t *testing.T) {
	var tests = []*struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxAreaV2(test.height))
	}
}
