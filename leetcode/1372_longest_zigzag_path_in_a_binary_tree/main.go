package main

func main() {
	// TODO: implement here
}

// refer to: https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/solutions/3432989/dfs-solution-with-explanation-c-go/?envType=study-plan-v2&envId=leetcode-75
func longestZigZagV1(root *TreeNode) int {
	return max(dfsV1(root.Left, true, 0), dfsV1(root.Right, false, 0))
}

func dfsV1(node *TreeNode, fromLeft bool, sum int) int {
	if node == nil {
		return sum
	}

	if fromLeft {
		// if previous node goes from left
		//   1. next node goes to right: sum + 1
		//   2. next node goes to left: reset sum to be zero
		return max(dfsV1(node.Left, true, 0), dfsV1(node.Right, false, sum+1))
	}
	// from right
	// if previous node goes from right
	//   1. next node goes to right: reset sum to be zero
	//   2. next node goes to left: sum + 1
	return max(dfsV1(node.Left, true, sum+1), dfsV1(node.Right, false, 0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
