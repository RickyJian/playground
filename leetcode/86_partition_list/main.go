package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// first: 第一個節點，後續回傳使用
	// firstAfterX: 第一個 val >= x 的節點，用來連結 val < x 最後的節點
	first, firstAfterX := &ListNode{}, &ListNode{}
	// beforeX: 保存所有 val < x 的 linked list
	// afterX: 保存所有 val >= x 的 linked list
	beforeX, afterX := first, firstAfterX
	for head != nil {
		if head.Val >= x {
			afterX.Next = head
			afterX = head
		} else {
			// head.Val < x
			beforeX.Next = head
			beforeX = beforeX.Next
		}
		head = head.Next
	}
	// 將最後一個節點指向 nil
	afterX.Next = nil
	// 連結最後一個 beforeX 的節點與第一個 afterX
	beforeX.Next = firstAfterX.Next
	return first.Next
}
