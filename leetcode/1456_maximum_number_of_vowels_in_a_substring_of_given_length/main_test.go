package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowelsV1(t *testing.T) {
	var tests = []*struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
		{
			s:        "weallloveyou",
			k:        7,
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxVowelsV1(test.s, test.k), test.s)
	}
}

func TestMaxVowelsV2(t *testing.T) {
	var tests = []*struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
		{
			s:        "weallloveyou",
			k:        7,
			expected: 4,
		},
		{
			s:        "ramadan",
			k:        2,
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxVowelsV2(test.s, test.k), test.s)
	}
}

func TestMaxVowelsV3(t *testing.T) {
	var tests = []*struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
		{
			s:        "weallloveyou",
			k:        7,
			expected: 4,
		},
		{
			s:        "ramadan",
			k:        2,
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxVowelsV3(test.s, test.k), test.s)
	}
}
