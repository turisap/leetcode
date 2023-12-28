package main

var testCases = []struct {
	name   string
	tree   *TreeNode
	result bool
}{
	{
		name: "[1,2,2,3,4,4,3]",
		tree: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
		result: true,
	},
	{
		name: "[1,2,2,3,4,4,8]",
		tree: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
				Left: &TreeNode{
					Val: 8,
				},
			},
		},
		result: false,
	}, {
		name: "[1,2,2,null,3,null,3]",
		tree: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		result: false,
	},
}
