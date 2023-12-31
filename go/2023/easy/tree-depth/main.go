package main

import "container/list"

func trdRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var inner func(node *TreeNode) int

	inner = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 1
		}

		var maxD int
		var lMD int
		var rMD int

		if node.Left != nil {
			lMD = inner(node.Left)
		}
		if node.Right != nil {
			rMD = inner(node.Right)
		}

		if rMD > lMD {
			maxD = rMD
		} else {
			maxD = lMD
		}

		return 1 + maxD
	}

	return inner(root)
}

func trdIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		n := queue.Len()
		for i := 0; i < n; i++ {
			el := queue.Front()
			queue.Remove(el)

			node := el.Value.(*TreeNode)

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
		maxDepth++
	}
	return maxDepth
}
