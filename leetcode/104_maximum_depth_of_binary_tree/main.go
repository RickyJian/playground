package main

func main() {
	// TODO: implement here
}

func maxDepthV1(root *TreeNode) int {
	return dfsV1(0, root)
}

func dfsV1(sum int, root *TreeNode) int {
	if root == nil {
		return sum
	}

	sum++
	// left
	left := dfsV1(sum, root.Left)
	// right
	right := dfsV1(sum, root.Right)
	var total int
	if left < right {
		total += right
	} else {
		total += left
	}
	return total
}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthV2(root.Left), maxDepthV2(root.Right)) + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
