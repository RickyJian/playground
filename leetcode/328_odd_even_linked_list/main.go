package main

func main() {
	// TODO: implement here
}

// refer to: https://leetcode.com/problems/odd-even-linked-list/solutions/78079/simple-o-n-time-o-1-space-java-solution/?envType=study-plan-v2&envId=leetcode-75
func oddEvenListV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	// keep even head after loop
	evenHead := even
	// if nodes are odd, when even is nil represents loop to the end
	// if node are even, when even next is null represents loop to the end
	// according to above rules are prevent odd or even invoice next attribute to meet nil pointer dereference
	for even != nil && even.Next != nil {
		// link odd node
		odd.Next = odd.Next.Next
		// link even node
		even.Next = even.Next.Next
		// go to next node
		odd = odd.Next
		even = even.Next
	}
	// link tail odd node and head even node
	odd.Next = evenHead
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
