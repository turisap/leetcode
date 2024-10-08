package max_depth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left) + 1
	rd := maxDepth(root.Right) + 1

	return int(math.Max(float64(ld), float64(rd)))
}
