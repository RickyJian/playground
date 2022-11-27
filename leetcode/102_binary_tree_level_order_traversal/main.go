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
	fmt.Println(levelOrderDFS(root))
	fmt.Println(levelOrderBFS(root))
}

func levelOrderDFS(root *TreeNode) [][]int {
	results := make([][]int, 0)
	dfs(root, 0, &results)
	return results
}

func dfs(root *TreeNode, level int, results *[][]int) {
	if root == nil {
		return
	} else if len(*results) == level {
		*results = append(*results, []int{})
	}

	(*results)[level] = append((*results)[level], root.Val)
	dfs(root.Left, level+1, results)
	dfs(root.Right, level+1, results)
}

func levelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var results [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var result []int
		for _, node := range queue {
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			result = append(result, node.Val)
		}
		results = append(results, result)
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
