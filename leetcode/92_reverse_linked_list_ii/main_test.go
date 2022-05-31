package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nodes    *ListNode
	left     int
	right    int
	expected *ListNode
}{
	{
		nodes: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
		left:  2,
		right: 5,
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	},
}

func TestReverseBetween(t *testing.T) {
	for _, test := range tests {
		results := reverseBetween(test.nodes, test.left, test.right)
		assert.EqualValues(t, test.expected, results)
	}
}
func TestReverseBetween2(t *testing.T) {
	for _, test := range tests {
		results := reverseBetween2(test.nodes, test.left, test.right)
		assert.EqualValues(t, test.expected, results)
	}
}

func TestReverseBetweenRecursion(t *testing.T) {
	for _, test := range tests {
		results := reverseBetweenRecursion(test.nodes, test.left, test.right)
		assert.EqualValues(t, test.expected, results)
	}
}
