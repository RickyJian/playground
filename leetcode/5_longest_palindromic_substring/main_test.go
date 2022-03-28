package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	input    string
	expected []string
}{
	{
		input: "a",
		expected: []string{
			"a",
		},
	},
	{
		input: "ac",
		expected: []string{
			"a",
			"c",
		},
	},
	{
		input: "babad",
		expected: []string{
			"bab",
			"aba",
		},
	},
	{
		input: "cbbd",
		expected: []string{
			"bb",
		},
	},
	{
		input: "aaacba",
		expected: []string{
			"aaa",
		},
	},
	{
		input: "aacabdkacaa",
		expected: []string{
			"aca",
		},
	},
}

func TestLongestPalindromeBrute(t *testing.T) {
	for _, test := range tests {
		result := longestPalindromeBrute(test.input)
		var existed bool
		for _, expected := range test.expected {
			if expected == result {
				existed = true
				break
			}
		}
		assert.Equal(t, true, existed)
	}
}

func TestLongestPalindromeDP(t *testing.T) {
	for _, test := range tests {
		result := longestPalindromeDP(test.input)
		var existed bool
		for _, expected := range test.expected {
			if expected == result {
				existed = true
				break
			}
		}
		assert.Equal(t, true, existed)
	}
}
