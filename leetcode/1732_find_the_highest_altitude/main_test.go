package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	var tests = []*struct {
		gain     []int
		expected int
	}{
		{
			gain:     []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			gain:     []int{-4, -3, -2, -1, 4, 3, 2},
			expected: 0,
		},
	}
	for _, test := range tests {
		assert.EqualValues(t, test.expected, largestAltitudeV1(test.gain), test.gain)
	}
}
