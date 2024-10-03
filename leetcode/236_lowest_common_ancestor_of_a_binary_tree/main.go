package main

func main() {
	// TODO: implement here
}

func lowestCommonAncestorV1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// use postorder to find children nodes are equal to p, or q
	left := lowestCommonAncestorV1(root.Left, p, q)
	right := lowestCommonAncestorV1(root.Right, p, q)
	if left != nil && right != nil {
		// if left and right both are not null, then current node is `LCA`
		return root
	} else if left != nil {
		// if left is not null, then
		return left
	}
	return right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
