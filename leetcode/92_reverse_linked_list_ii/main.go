package main

func main() {
	results := reverseBetween(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}, 3, 4)
	for results != nil {
		results = results.Next
	}
	results2 := reverseBetween2(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}, 1, 4)
	for results != nil {
		results = results.Next
	}
	for results2 != nil {
		results2 = results2.Next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	current := head
	var first, second, prev *ListNode
	for count := 0; count < right; count++ {
		nextNode := head.Next
		if count == left-1 {
			first = prev
			second = head
		} else if count == right-1 {
			if first == nil {
				current = head
			} else {
				// first existed
				first.Next = head
			}
			second.Next = head.Next
			head.Next = prev
		} else if count > left-1 {
			head.Next = prev
		}
		prev = head
		head = nextNode
	}
	return current
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	//
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	start := prev.Next
	next := start.Next
	for i := 0; i < right-left; i++ {
		// 跟下一個 node 做交換
		start.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = start.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
