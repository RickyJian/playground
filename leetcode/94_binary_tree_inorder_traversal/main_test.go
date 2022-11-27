package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	root     *TreeNode
	expected []int
}{
	{
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		expected: []int{1, 2, 3},
	},
}

func TestPreorderTraversalDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, inorderTraversalDFS(test.root))
	}
}

func TestPreorderTraversalIterator(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, inorderTraversalIterator(test.root))
	}
}
