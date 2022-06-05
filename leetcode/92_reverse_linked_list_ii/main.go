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
	// 每次 loop 將時
	//   1. 將 start node 與 next node 位置交換
	//     i. start.Next 設定為 next.Next
	//    ii. next.Next 設定為 prev.Next (prev.Next = start node)
	//   2. 調整 prev node 的指向， prev.Next = next node
	//   3. 向下個 node 移動
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

// refer to: https://leetcode.com/problems/reverse-linked-list-ii/discuss/30673/Why-nobody-does-it-with-recursion-shouldn't-the-code-be-simpler
func reverseBetweenRecursion(head *ListNode, left int, right int) *ListNode {
	if left == right {
		// 節點走訪到最後一個要反轉的位置
		return head
	}

	if left > 1 {
		// 尚未到需要反轉的節點，持續走訪下個節點
		//   ＊ -1：走訪過的節點
		head.Next = reverseBetweenRecursion(head.Next, left-1, right-1)
		return head
	} else {
		// left == 1 時，代表 head 到達了我們要調換節點的位置
		next := head.Next
		// revereNode 皆為最後節點
		reverseNode := reverseBetweenRecursion(next, 1, right-1)
		// 調換節點
		// head 節點改指向 next 下個節點位置
		head.Next = next.Next
		// next 下個節點改指向 head 節點
		next.Next = head
		// 由於節點反轉的關係，最後一個節點須移到 left 的位置
		return reverseNode
	}
}

// TODO: two pointer swap

type ListNode struct {
	Val  int
	Next *ListNode
}
