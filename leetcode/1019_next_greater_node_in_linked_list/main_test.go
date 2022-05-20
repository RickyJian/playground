package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data1 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}
	data2 = &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}
)

func TestNextLargerNodesStack1(t *testing.T) {
	var tests = []*struct {
		listNode *ListNode
		expected []int
	}{
		{
			listNode: data1,
			expected: []int{5, 5, 0},
		},
		{
			listNode: data2,
			expected: []int{7, 0, 5, 5, 0},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, nextLargerNodesStack1(test.listNode))
	}
}

func TestNextLargerNodesStack2(t *testing.T) {
	var tests = []*struct {
		listNode *ListNode
		expected []int
	}{
		{
			listNode: data1,
			expected: []int{5, 5, 0},
		},
		{
			listNode: data2,
			expected: []int{7, 0, 5, 5, 0},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, nextLargerNodesStack2(test.listNode))
	}
}

func TestNextLargerNodesStack3(t *testing.T) {
	var tests = []*struct {
		listNode *ListNode
		expected []int
	}{
		{
			listNode: data1,
			expected: []int{5, 5, 0},
		},
		{
			listNode: data2,
			expected: []int{7, 0, 5, 5, 0},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, nextLargerNodesStack3(test.listNode))
	}
}
