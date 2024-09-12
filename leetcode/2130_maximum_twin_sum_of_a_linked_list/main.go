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

type ListNode struct {
	Val  int
	Next *ListNode
}
