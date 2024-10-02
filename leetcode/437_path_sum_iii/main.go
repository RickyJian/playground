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

// prefix sum solution
// refer to: https://leetcode.com/problems/path-sum-iii/solutions/141424/python-step-by-step-walk-through-easy-to-understand-two-solutions-comparison/?envType=study-plan-v2&envId=leetcode-75
func pathSumV3(root *TreeNode, targetSum int) int {
	return dfsV3(root, targetSum, 0, make(map[int]int))
}

func dfsV3(node *TreeNode, target, sum int, cache map[int]int) int {
	if node == nil {
		return 0
	}

	var count int
	sum += node.Val
	if target == sum {
		count++
	}
	// 倘若 `sum-target` 出現在 cache 中， 代表父節點到當下節點的總和 = target。
	count += cache[sum-target]
	// 使用 cache 儲存 sum 出現的次數，用來查找先前路徑的總和。
	cache[sum]++
	count += dfsV3(node.Left, target, sum, cache)
	count += dfsV3(node.Right, target, sum, cache)
	// 由於無後續節點需要走訪，因此將先前加總過得次數移除。
	cache[sum]--
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
