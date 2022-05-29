package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nodes    *ListNode
	expected *ListNode
}{
	{
		nodes: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		expected: &ListNode{
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
}

func TestReverseList(t *testing.T) {
	for _, test := range tests {
		results := reverseList(test.nodes)
		assert.EqualValues(t, test.expected, results)
	}
}

func TestReverseListSpaceEnhance(t *testing.T) {
	for _, test := range tests {
		results := reverseListSpaceEnhance(test.nodes)
		assert.EqualValues(t, test.expected, results)
	}
}

func TestReverseListRecursion(t *testing.T) {
	for _, test := range tests {
		results := reverseListRecursion(test.nodes)
		assert.EqualValues(t, test.expected, results)
	}
}
