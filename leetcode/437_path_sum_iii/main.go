package main

func main() {
	// TODO: implement here
}

func pathSumV1(root *TreeNode, targetSum int) int {
	return dfsV1(root, targetSum, make([]int, 0))
}

func dfsV1(node *TreeNode, target int, prevNodes []int) int {
	if node == nil {
		return 0
	}

	var count int
	if target == node.Val {
		count++
	}
	for i := range prevNodes {
		prevNodes[i] += node.Val
		if node.Val == 1 {
		}
		if prevNodes[i] == target {
			count++
		}
	}
	left := make([]int, len(prevNodes)+1)
	copy(left, prevNodes)
	left[len(left)-1] = node.Val
	right := make([]int, len(left))
	copy(right, left)
	count += dfsV1(node.Left, target, left)
	count += dfsV1(node.Right, target, right)
	return count
}

func pathSumV2(root *TreeNode, targetSum int) int {
	return dfsV2(root, targetSum, make([]int, 0))
}

func dfsV2(node *TreeNode, target int, prevNodes []int) int {
	if node == nil {
		return 0
	}

	nextNodes := make([]int, len(prevNodes)+1)
	copy(nextNodes, prevNodes)
	nextNodes[len(nextNodes)-1] = node.Val

	var count, sum int
	for i := len(nextNodes) - 1; i >= 0; i-- {
		// 愈後面的元素代表遇愈底層的節點，
		// 從最後加總到第一個數值代表 parent 到 child 的總和。
		sum += nextNodes[i]
		if sum == target {
			count++
		}
	}
	count += dfsV2(node.Left, target, nextNodes)
	count += dfsV2(node.Right, target, nextNodes)
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
