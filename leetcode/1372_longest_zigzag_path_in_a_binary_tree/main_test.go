package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestZigZagV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: 1,
								},
							},
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
							},
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			expected: 4,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, longestZigZagV1(test.root))
	}
}
