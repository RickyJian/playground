package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nodes    *ListNode
	x        int
	expected *ListNode
}{
	{
		nodes: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
		},
		x: 3,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	},
	{
		nodes: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		x: 4,
		expected: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	},
	{
		nodes: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		x: 1,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	},
}

func TestPartition(t *testing.T) {
	for _, test := range tests {
		assert.EqualValues(t, test.expected, partition(test.nodes, test.x))
	}
}
