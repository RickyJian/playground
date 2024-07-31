package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifferenceV1(t *testing.T) {
	var tests = []*struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
			expected: [][]int{
				{1, 3},
				{4, 6},
			},
		},
		{
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			expected: [][]int{
				{3},
				{},
			},
		},
	}
	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, findDifferenceV1(test.nums1, test.nums2), test.nums1)
	}
}

func TestFindDifferenceV2(t *testing.T) {
	var tests = []*struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		// {
		// 	nums1: []int{1, 2, 3},
		// 	nums2: []int{2, 4, 6},
		// 	expected: [][]int{
		// 		{1, 3},
		// 		{4, 6},
		// 	},
		// },
		{
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			expected: [][]int{
				{3},
				{},
			},
		},
	}
	for _, test := range tests {
		assert.ElementsMatch(t, test.expected, findDifferenceV2(test.nums1, test.nums2), test.nums1)
	}
}
