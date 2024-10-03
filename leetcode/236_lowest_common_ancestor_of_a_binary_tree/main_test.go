package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestorV1(t *testing.T) {
	var tests = []*struct {
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 1},
			expected: &TreeNode{Val: 3},
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 4},
			expected: &TreeNode{Val: 5},
		},
		// {
		// 	root: &TreeNode{
		// 		Val: 1,
		// 		Left: &TreeNode{
		// 			Val: 2,
		// 		},
		// 	},
		// 	p:        &TreeNode{Val: 1},
		// 	q:        &TreeNode{Val: 2},
		// 	expected: &TreeNode{Val: 1},
		// },
	}
	for _, test := range tests {
		assert.Equal(t, test.expected.Val, lowestCommonAncestorV1(test.root, test.p, test.q).Val)
	}
}
