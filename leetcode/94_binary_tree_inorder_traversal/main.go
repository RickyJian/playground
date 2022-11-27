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
	fmt.Println(inorderTraversalDFS(root))
	// fmt.Println(inorderTraversalIterator(root))
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
	// TODO: implement
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
