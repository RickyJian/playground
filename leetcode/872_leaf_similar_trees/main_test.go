package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafSimilarV1(t *testing.T) {
	var tests = []*struct {
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		{
			root1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			root2: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			root1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			root2: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, leafSimilarV1(test.root1, test.root2))
	}
}
