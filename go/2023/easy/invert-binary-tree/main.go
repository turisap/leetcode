package main

func ibtRecursive(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = ibtRecursive(root.Right), ibtRecursive(root.Left)
	}

	return root
}

func ibtIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		curr := q[0]

		tmp := curr.Left
		curr.Left = curr.Right
		curr.Right = tmp

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		q = q[1:]
	}

	return root
}
