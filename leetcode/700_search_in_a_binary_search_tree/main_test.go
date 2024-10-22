package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBSTV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		val      int
		expected *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val: 2,
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, searchBSTV1(test.root, test.val))
	}
}
