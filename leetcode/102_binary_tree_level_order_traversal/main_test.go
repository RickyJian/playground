package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	root     *TreeNode
	expected [][]int
}{
	{
		root: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Left: &TreeNode{
				Val: 9,
			},
		},
		expected: [][]int{
			{3},
			{9, 20},
			{15, 7},
		},
	},
}

func TestLevelOrderDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, levelOrderDFS(test.root), test.expected)
	}
}

func TestLevelOrderBFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, levelOrderBFS(test.root), test.expected)
	}
}
