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
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
