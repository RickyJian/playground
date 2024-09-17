package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepthV1(t *testing.T) {
	var tests = []*struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
			},
			expected: 3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxDepthV1(test.input), test.input.Val)
	}
}

func TestMaxDepthV2(t *testing.T) {
	var tests = []*struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
			},
			expected: 3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxDepthV2(test.input), test.input.Val)
	}
}
