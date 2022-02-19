package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonadeChange(t *testing.T) {
	var tests = []*struct {
		bills    []int
		expected bool
	}{
		{
			bills:    []int{5, 5, 5, 10, 20},
			expected: true,
		},
		{
			bills:    []int{5, 5, 10, 10, 20},
			expected: false,
		},
		{
			bills:    []int{5, 5, 5, 5, 10, 5, 10, 10, 10, 20},
			expected: true,
		},
	}
	for _, test := range tests {
		results := lemonadeChange(test.bills)
		assert.Equal(t, test.expected, results)
	}
}
