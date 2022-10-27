package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	s        string
	t        string
	expected int
}{
	{
		s:        "aba",
		t:        "baba",
		expected: 6,
	},
	{
		s:        "ab",
		t:        "bb",
		expected: 3,
	},
}

func TestCountSubstringsN5(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, countSubstringsN5(test.s, test.t))
	}
}

func TestCountSubstringsN3(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, countSubstringsN3(test.s, test.t))
	}
}
