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
	// 由於反轉的關係第一個 node 指向可設為 nil
	current.Next = nil
	for next != nil {
		// 保存下下個 node 為後續走訪節點做準備
		nextNode := next.Next
		// 下一個 node 指向當下的 node
		next.Next = current
		// 走訪下個 node
		//   1. 將下個 node 設定為當下 node
		//   2. 下個 node 為原本的 node 的指向
		current = next
		next = nextNode
	}
	return current
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 最後一個 node，也就是 reverse 的頭
	nextNode := reverseListRecursion(head.Next)
	// 將下一個 node 指向當下的 node
	// 若為傳進 function 最後一個 node，head.Next = nextNode
	head.Next.Next = head
	// 當 node 為傳進 function 的第一個 node 時將他設為 nil
	head.Next = nil
	return nextNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
