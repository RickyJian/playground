package main

import "slices"

func main() {
	// TODO: implement here
}

func leafSimilarV1(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(dfs(root1), dfs(root2))
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(dfs(root.Left), dfs(root.Right)...)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
