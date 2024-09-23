package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodesV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 4,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 3,
			},
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, goodNodesV1(test.root), test.root.Val)
	}
}
