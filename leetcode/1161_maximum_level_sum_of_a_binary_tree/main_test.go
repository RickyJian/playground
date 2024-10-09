package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLevelSumV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxLevelSumV1(test.root))
	}
}
