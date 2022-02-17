package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicateV1(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			k:        0,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        1,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        2,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        4,
			expected: true,
		},
	}
	for _, test := range tests {
		results := containsNearbyDuplicateV1(test.nums, test.k)
		assert.Equal(t, test.expected, results)
	}
}

func TestContainsNearbyDuplicateV2(t *testing.T) {
	var tests = []*struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			k:        0,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        1,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        2,
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1},
			k:        4,
			expected: true,
		},
	}
	for _, test := range tests {
		results := containsNearbyDuplicateV2(test.nums, test.k)
		assert.Equal(t, test.expected, results)
	}
}
