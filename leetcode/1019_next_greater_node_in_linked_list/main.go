package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextLargerNodesStack2(&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}))
	fmt.Println(nextLargerNodesStack2(
		&ListNode{Val: 2,
			Next: &ListNode{Val: 7,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 5}}}}}))
}

func nextLargerNodesStack1(head *ListNode) []int {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	results := make([]int, len(values))
	var stack []int
	for i, val := range values {
		for len(stack) > 0 && val > values[stack[len(stack)-1]] {
			results[stack[len(stack)-1]] = val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return results
}

func nextLargerNodesStack2(head *ListNode) []int {
	var stack, values []int
	results := make([]int, 10000)
	for i := 0; head != nil; i++ {
		values = append(values, head.Val)
		for len(stack) > 0 && head.Val > values[stack[len(stack)-1]] {
			results[stack[len(stack)-1]] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		head = head.Next
	}
	return results[:len(values)]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
