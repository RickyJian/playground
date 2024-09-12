package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenListV1(t *testing.T) {
	var tests = []*struct {
		head     *ListNode
		expected int
	}{
		{
			head: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			expected: 6,
		},
		{
			head: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			expected: 7,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, pairSumV1(test.head), test.head.Val)
	}
}
