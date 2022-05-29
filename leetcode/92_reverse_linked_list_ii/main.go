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

	// 將原先的 linked list 源頭增加一個節點，
	// 預防從一第一個 node 開始時的判斷
	dummy := &ListNode{Next: head}
	prev := dummy
	// 取到 left 前的節點
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	start := prev.Next
	next := start.Next
	// 每次 loop 將
	//   1. 當下 node(start) 向後移，
	//   2. 下個 node(next) 向前移
	//   3. 前一個 node(prev) 指向下個 node(next)
	//   4. 往當下的下個 node(start.Next) 移動
	for i := 0; i < right-left; i++ {
		// 當下 node 向後移
		start.Next = next.Next
		// 下個 node 向前移
		next.Next = prev.Next
		// 前一個 node 指向下個 node
		prev.Next = next
		// 往下個 node(start.Next) 移動
		next = start.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
