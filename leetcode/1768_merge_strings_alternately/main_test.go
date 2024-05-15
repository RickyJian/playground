package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	var tests = []*struct {
		word1    string
		word2    string
		expected string
	}{
		{
			word1:    "abc",
			word2:    "pqr",
			expected: "apbqcr",
		},
		{
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
		{
			word1:    "abcd",
			word2:    "pq",
			expected: "apbqcd",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, mergeAlternately(test.word1, test.word2))
	}
}
