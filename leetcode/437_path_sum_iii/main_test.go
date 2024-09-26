package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSumV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		target   int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: -2,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			target:   8,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			target:   1,
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, pathSumV1(test.root, test.target))
	}
}

func TestPathSumV2(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		target   int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: -2,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			target:   8,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			target:   1,
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, pathSumV2(test.root, test.target))
	}
}
