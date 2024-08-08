package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	var tests = []*struct {
		word1    string
		word2    string
		expected bool
	}{
		{
			word1:    "abc",
			word2:    "cba",
			expected: true,
		},
		{
			word1:    "a",
			word2:    "aa",
			expected: false,
		},
		{
			word1:    "cabbba",
			word2:    "abbccc",
			expected: true,
		},
		{
			word1:    "cabbba",
			word2:    "aabbss",
			expected: false,
		},
		{
			word1:    "abbzzca",
			word2:    "babzzcz",
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, closeStringsV1(test.word1, test.word2), test.word1)
	}
}
