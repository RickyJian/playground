package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressV1(t *testing.T) {
	var tests = []*struct {
		chars          []byte
		expected       int
		expectedCharts []byte
	}{
		{
			chars:          []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected:       6,
			expectedCharts: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			chars:          []byte{'a'},
			expected:       1,
			expectedCharts: []byte{'a'},
		},
		{
			chars:          []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected:       4,
			expectedCharts: []byte{'a', 'b', '1', '2'},
		},
		{
			chars:          []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			expected:       6,
			expectedCharts: []byte{'a', '3', 'b', '2', 'a', '2'},
		},
	}
	for _, test := range tests {
		q := make([]byte, len(test.chars))
		copy(q, test.chars)
		result := compressV1(test.chars)
		assert.Equal(t, test.expected, result, q)
		assert.Equal(t, test.expectedCharts, test.chars, q)
	}
}
