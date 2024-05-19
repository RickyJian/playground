package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDOfStrings(t *testing.T) {
	var tests = []*struct {
		str1     string
		str2     string
		expected string
	}{
		{
			str1:     "ABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			str1:     "ABABAB",
			str2:     "ABAB",
			expected: "AB",
		},
		{
			str1:     "LEET",
			str2:     "CODE",
			expected: "",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, gcdOfStrings(test.str1, test.str2))
	}
}
