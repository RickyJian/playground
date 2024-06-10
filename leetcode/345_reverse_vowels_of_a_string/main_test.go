package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	var tests = []*struct {
		s        string
		expected string
	}{
		{
			s:        "hello",
			expected: "holle",
		},
		{
			s:        "leetcode",
			expected: "leotcede",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, reverseVowels(test.s))
	}
}
