package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	var tests = []*struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
			expected:  true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, canPlaceFlowers(test.flowerbed, test.n))
	}
}
