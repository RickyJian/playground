package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	s        string
	wordDict []string
	expected bool
}{
	{
		s:        "leetcode",
		wordDict: []string{"leet", "code"},
		expected: true,
	},
	{
		s:        "applepenapple",
		wordDict: []string{"apple", "pen"},
		expected: true,
	},
	{
		s:        "catsandog",
		wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		expected: false,
	},
}

func TestWordBreak(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, wordBreak(test.s, test.wordDict))
	}
}

func TestWordBreakDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, wordBreakDFS(test.s, test.wordDict))
	}
}
