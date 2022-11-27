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
		expected: []int{3, 2, 1},
	},
}

func TestPostorderTraversalDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, postorderTraversalDFS(test.root))
	}
}

func TestPostorderTraversalIterator(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, postorderTraversalIterator(test.root))
	}
}
