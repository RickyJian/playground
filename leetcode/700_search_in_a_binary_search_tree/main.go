package main

func main() {
	// TODO: implement here
}

func searchBSTV1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBSTV1(root.Left, val)
	}
	return searchBSTV1(root.Right, val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
