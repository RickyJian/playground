package main

import "math"

func main() {
	// TODO: implement here
}

func maxLevelSumV1(root *TreeNode) int {
	maxSum := math.MinInt
	var result, level int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level++
		size := len(queue)
		var sum int
		for i := 0; i < size; i++ {
			top := queue[0]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			sum += top.Val
			queue = queue[1:]
		}
		if maxSum < sum {
			maxSum = sum
			result = level
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
