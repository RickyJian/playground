package main

func main() {
	// TODO: implement here
}

func pairSumV1(head *ListNode) int {
	var arrs []int
	for {
		if head == nil {
			break
		}
		arrs = append(arrs, head.Val)
		head = head.Next
	}
	var max int
	for i, j := 0, len(arrs)-1; i < len(arrs)/2; i, j = i+1, j-1 {
		sum := arrs[i] + arrs[j]
		if sum > max {
			max = sum
		}
	}
	return max
}

func pairSumV2(head *ListNode) int {
	// find middle node
	slow := head
	fast := head.Next
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	// keep after middle node
	right := slow.Next
	// remove link to next node, due to we need to reverse left nodes
	slow.Next = nil
	// reverse slow
	left := reverse(nil, head)

	// sum twin
	var max int
	for {
		if right == nil {
			break
		}
		sum := left.Val + right.Val
		if sum > max {
			max = sum
		}
		left = left.Next
		right = right.Next
	}
	return max
}

func reverse(prev, head *ListNode) *ListNode {
	if head == nil {
		return prev
	}

	next := head.Next
	head.Next = prev
	return reverse(head, next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
