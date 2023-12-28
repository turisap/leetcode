package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btiRecursive(root *TreeNode) []int {
	res := []int{}

	var inner func(root *TreeNode)

	inner = func(r *TreeNode) {

		if r == nil {
			return
		}

		if r.Left != nil {
			inner(r.Left)
		}

		res = append(res, r.Val)

		if r.Right != nil {
			inner(r.Right)
		}
	}

	inner(root)

	return res
}

func btiIterative(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		l := len(stack)
		root = stack[l-1]
		stack = stack[:l-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return res
}
