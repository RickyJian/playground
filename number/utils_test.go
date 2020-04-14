package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitCount1(t *testing.T) {
	tests := []struct {
		no       int
		expected int
	}{
		{
			no:       2,
			expected: 1,
		},
		{
			no:       12,
			expected: 2,
		},
		{
			no:       112,
			expected: 3,
		},
		{
			no:       1122,
			expected: 4,
		},
		{
			no:       1,
			expected: 1,
		},
		{
			no:       10,
			expected: 2,
		},
		{
			no:       100,
			expected: 3,
		},
		{
			no:       1000,
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, DigitCount1(test.no), test.expected)
	}
}

func TestDigitCount2(t *testing.T) {
	tests := []struct {
		no       int
		expected int
	}{
		{
			no:       2,
			expected: 1,
		},
		{
			no:       12,
			expected: 2,
		},
		{
			no:       112,
			expected: 3,
		},
		{
			no:       1122,
			expected: 4,
		},
		{
			no:       1,
			expected: 1,
		},
		{
			no:       10,
			expected: 2,
		},
		{
			no:       100,
			expected: 3,
		},
		{
			no:       1000,
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, DigitCount2(test.no), test.expected)
	}
}
