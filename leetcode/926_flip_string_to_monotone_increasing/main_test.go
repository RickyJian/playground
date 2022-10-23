package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFlipsMonoIncrBrute(t *testing.T) {
	var tests = []*struct {
		str      string
		expected int
	}{
		{
			str:      "00110",
			expected: 1,
		},
		{
			str:      "010110",
			expected: 2,
		},
		{
			str:      "00011000",
			expected: 2,
		},
		{
			str:      "11011",
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, minFlipsMonoIncrBrute(test.str))
	}
}

func TestMinFlipsMonoIncrDP2Array(t *testing.T) {
	var tests = []*struct {
		str      string
		expected int
	}{
		{
			str:      "00110",
			expected: 1,
		},
		{
			str:      "010110",
			expected: 2,
		},
		{
			str:      "00011000",
			expected: 2,
		},
		{
			str:      "11011",
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, minFlipsMonoIncrDP2Array(test.str))
	}
}

func TestMinFlipsMonoIncrDP1Array(t *testing.T) {
	var tests = []*struct {
		str      string
		expected int
	}{
		{
			str:      "00110",
			expected: 1,
		},
		{
			str:      "010110",
			expected: 2,
		},
		{
			str:      "00011000",
			expected: 2,
		},
		{
			str:      "11011",
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, minFlipsMonoIncrDP1Array(test.str))
	}
}
