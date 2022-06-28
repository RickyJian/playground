package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	n        int
	k        int
	expected byte
}{
	{
		n:        4,
		k:        1,
		expected: '0',
	},
	{
		n:        4,
		k:        2,
		expected: '1',
	},
	{
		n:        4,
		k:        4,
		expected: '1',
	},
	{
		n:        4,
		k:        5,
		expected: '0',
	},
	{
		n:        4,
		k:        8,
		expected: '1',
	},
	{
		n:        4,
		k:        13,
		expected: '0',
	},
	{
		n:        4,
		k:        15,
		expected: '1',
	},
}

func TestFindKthBit(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, findKthBit(test.n, test.k))
	}
}
