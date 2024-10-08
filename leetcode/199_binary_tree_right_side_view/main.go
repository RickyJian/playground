package main

func main() {
	// TODO: implement here
}

func rightSideViewV1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := make([]int, 0)
	bfsV1([]*TreeNode{root}, 0, &results)
	return results
}

func bfsV1(nodes []*TreeNode, level int32, results *[]int) {
	if len(nodes) == 0 {
		return
	} else if len(nodes) > 0 {
		*results = append(*results, nodes[len(nodes)-1].Val)
	}
	var queue []*TreeNode
	for len(nodes) > 0 {
		top := nodes[0]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
		nodes = nodes[1:]
	}
	bfsV1(queue, level+1, results)
}

func rightSideViewV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var results []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			if i == size-1 {
				results = append(results, top.Val)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			queue = queue[1:]
		}
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
