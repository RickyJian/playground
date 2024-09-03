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

type ListNode struct {
	Val  int
	Next *ListNode
}
