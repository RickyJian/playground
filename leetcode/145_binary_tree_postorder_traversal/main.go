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
	fmt.Println(postorderTraversalDFS(root))
	fmt.Println(postorderTraversalIterator(root))
}

func postorderTraversalDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var results []int
	results = append(results, postorderTraversalDFS(root.Left)...)
	results = append(results, postorderTraversalDFS(root.Right)...)
	results = append(results, root.Val)
	return results
}

func postorderTraversalIterator(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var results []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left == nil && cur.Right == nil {
			results = append(results, cur.Val)
			continue
		}

		stack = append(stack, cur)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		cur.Left, cur.Right = nil, nil
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
