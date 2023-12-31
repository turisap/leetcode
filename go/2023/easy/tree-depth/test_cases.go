package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var testCases = []struct {
	name   string
	input  *TreeNode
	result int
}{
	{
		"[3,9,20,null,null,15,7]",
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		3,
	},
	{
		"[1,nil,3]",
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
			},
		},
		2,
	},
	{
		"[1,2,3]",
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
			Left: &TreeNode{
				Val: 3,
			},
		},
		2,
	},
	{
		"[1]",
		&TreeNode{
			Val: 1,
		},
		1,
	},
	{
		"[]",
		nil,
		0,
	},
}
