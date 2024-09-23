package main

func main() {
	// TODO: implement here
}

func goodNodesV1(root *TreeNode) int {
	return dfsV1(root.Val, root)
}

func dfsV1(prev int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	current := root.Val
	var count int
	if current >= prev {
		count++
		prev = current
	}
	count += dfsV1(prev, root.Left)
	count += dfsV1(prev, root.Right)
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
