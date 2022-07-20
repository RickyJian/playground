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
		s:        " ",
		wordDict: []string{"leet", "code"},
		expected: false,
	},
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
	{
		s:        "catsanddog",
		wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		expected: true,
	},
	{
		s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		expected: false,
	},
}

func TestWordBreakDP(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, wordBreakDP(test.s, test.wordDict))
	}
}

func TestWordBreakDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, wordBreakDFS(test.s, test.wordDict))
	}
}

func TestWordBreakDFS2(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, wordBreakDFS2(test.s, test.wordDict))
	}
}
