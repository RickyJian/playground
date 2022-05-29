package main

func main() {
}

func reverseList(head *ListNode) *ListNode {
	var newNode *ListNode
	for head != nil {
		if newNode == nil {
			newNode = &ListNode{Val: head.Val}
		} else {
			newNode = &ListNode{Val: head.Val, Next: newNode}
		}
		head = head.Next
	}
	return newNode
}

func reverseListSpaceEnhance(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current, next := head, head.Next
	current.Next = nil
	for next != nil {
		nextNode := next.Next
		next.Next = current
		current = next
		next = nextNode
	}
	return current
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextNodes := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return nextNodes
}

type ListNode struct {
	Val  int
	Next *ListNode
}
