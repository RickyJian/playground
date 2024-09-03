package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddleV1(t *testing.T) {
	var tests = []*struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
			},
			expected: nil,
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			expected: &ListNode{
				Val: 2,
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, deleteMiddleV1(test.head), test.head.Val)
	}
}

func TestDeleteMiddleV2(t *testing.T) {
	var tests = []*struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
			},
			expected: nil,
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			expected: &ListNode{
				Val: 2,
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, deleteMiddleV2(test.head), test.head.Val)
	}
}

func TestDeleteMiddleV3(t *testing.T) {
	var tests = []*struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
			},
			expected: nil,
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			expected: &ListNode{
				Val: 2,
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, deleteMiddleV3(test.head), test.head.Val)
	}
}
