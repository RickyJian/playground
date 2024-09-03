package main

func main() {
	// TODO: implement here
}

func deleteMiddleV1(head *ListNode) *ListNode {
	top := head
	var nodes []*ListNode
	for {
		if head == nil {
			break
		}
		nodes = append(nodes, head)
		head = head.Next
	}
	nodeLen := len(nodes)
	middle := nodeLen / 2
	if middle == 0 {
		return nil
	} else if nodeLen > middle+1 {
		prev := nodes[len(nodes)/2-1]
		prev.Next = nodes[len(nodes)/2+1]
	} else {
		top.Next = nil
	}
	return top
}

func deleteMiddleV2(head *ListNode) *ListNode {
	top := head
	var count int
	for {
		if head == nil {
			break
		}
		count++
		head = head.Next
	}
	middle := count / 2
	if middle == 0 {
		return nil
	}

	result := top
	for i := 0; i < middle-1; i++ {
		if i < middle {
			top = top.Next
		}
	}
	top.Next = top.Next.Next
	return result
}

// refer to: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/solutions/4705282/simple-beginner-friendly-dry-run-two-pointer-approach-time-o-n-space-o-1-gits/
func deleteMiddleV3(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
