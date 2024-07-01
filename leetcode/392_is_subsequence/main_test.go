package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequenceV1(t *testing.T) {
	var tests = []*struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			s:        "",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "b",
			t:        "abc",
			expected: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, isSubsequenceV1(test.s, test.t))
	}
}
