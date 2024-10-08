package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideViewV1(t *testing.T) {
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
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, rightSideViewV1(test.root))
	}
}

func TestRightSideViewV2(t *testing.T) {
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
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, rightSideViewV2(test.root))
	}
}
