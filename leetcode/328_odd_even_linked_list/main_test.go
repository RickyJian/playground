package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenListV1(t *testing.T) {
	var tests = []*struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
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
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, oddEvenListV1(test.head), test.head.Val)
	}
}
