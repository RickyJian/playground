package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(preorderTraversalDFS(root))
	fmt.Println(preorderTraverseIterator(root))
}

// 中 -> 左 -> 右
func preorderTraversalDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	results := []int{root.Val}
	results = append(results, preorderTraversalDFS(root.Left)...)
	results = append(results, preorderTraversalDFS(root.Right)...)
	return results
}

func preorderTraverseIterator(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var results []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		results = append(results, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
