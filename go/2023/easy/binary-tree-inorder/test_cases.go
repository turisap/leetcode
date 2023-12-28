package main

var testCases = []struct {
	name   string
	tree   *TreeNode
	result string
}{
	{
		name: "[1,null,2,3]",
		tree: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		result: "[1 2 3]",
	},
	{
		name:   "empty",
		tree:   nil,
		result: "[]",
	},
	{
		name:   "root only",
		tree:   &TreeNode{Val: 1},
		result: "[1]",
	},
}
