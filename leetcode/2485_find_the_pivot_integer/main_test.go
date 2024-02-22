package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	n        int
	expected int
}{
	{
		n:        8,
		expected: 6,
	},
	{
		n:        1,
		expected: 1,
	},
	{
		n:        4,
		expected: -1,
	},
}

func TestPivotInteger(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, pivotIntegerV1(test.n))
	}
}
