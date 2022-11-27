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
	// fmt.Println(inorderTraversalDFS(root))
	fmt.Println(inorderTraversalIterator(root))
}

// 左 -> 中 -> 右
func inorderTraversalDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var results []int
	results = append(results, inorderTraversalDFS(root.Left)...)
	results = append(results, root.Val)
	results = append(results, inorderTraversalDFS(root.Right)...)
	return results
}

func inorderTraversalIterator(root *TreeNode) []int {
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

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		stack = append(stack, cur)
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
