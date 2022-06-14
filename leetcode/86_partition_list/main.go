package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	first, beforeAfterX := &ListNode{}, &ListNode{}
	beforeX, afterX := first, beforeAfterX
	for head != nil {
		if x > head.Val {
			beforeX.Next = head
			beforeX = beforeX.Next
		} else {
			afterX.Next = head
			afterX = head
		}
		head = head.Next
	}
	afterX.Next = nil
	beforeX.Next = beforeAfterX.Next
	return first.Next
}
