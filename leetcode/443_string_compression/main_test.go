package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressV1(t *testing.T) {
	var tests = []*struct {
		chars    []byte
		expected int
	}{
		{
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			chars:    []byte{'a'},
			expected: 1,
		},
		{
			chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
		{
			chars:    []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			expected: 6,
		},
	}
	for _, test := range tests {
		q := make([]byte, len(test.chars))
		copy(q, test.chars)
		result := compressV1(q)
		assert.Equal(t, test.expected, result, q)
	}
}
