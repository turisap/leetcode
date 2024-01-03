package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var testCases = []struct {
	name   string
	head   *TreeNode
	result *TreeNode
}{
	{
		"1-2-3",
		&TreeNode{Val: 2, Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 3}},
		&TreeNode{Val: 2, Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 3}},
	},
	{
		"complex",
		&TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			Left: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 6,
				},
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
		&TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 3,
				},
				Left: &TreeNode{
					Val: 1,
				},
			},
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 9,
				},
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
	},
}
