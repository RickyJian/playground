package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	s        string
	wordDict []string
	expected []string
}{
	{
		s:        "catsanddog",
		wordDict: []string{"cat", "cats", "and", "sand", "dog"},
		expected: []string{"cats and dog", "cat sand dog"},
	},
	{
		s:        "pineapplepenapple",
		wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
		expected: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
	},
	{
		s:        "catsandog",
		wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		expected: []string{},
	},
}

func TestWordBreak(t *testing.T) {
	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, wordBreak(test.s, test.wordDict))
	}
}

func TestWordBreakWithMap(t *testing.T) {
	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, wordBreakWithMap(test.s, test.wordDict))
	}
}
