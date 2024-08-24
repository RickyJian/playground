package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeStringV1(t *testing.T) {
	var tests = []*struct {
		s        string
		expected string
	}{
		{
			s:        "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			s:        "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			s:        "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, decodeStringV1(test.s), test.s)
	}
}
