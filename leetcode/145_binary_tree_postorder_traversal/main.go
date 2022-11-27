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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
