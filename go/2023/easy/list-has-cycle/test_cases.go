package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeCircledList() *ListNode {
	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 3, Next: node2}

	node4.Next = node2

	return head
}

var testCases = []struct {
	name   string
	head   *ListNode
	result bool
}{
	{
		name:   "true for [3,2,0,-4]",
		head:   makeCircledList(),
		result: true,
	},
	{
		name:   "false for [1]",
		head:   &ListNode{Val: 1},
		result: false,
	},
	{
		name:   "false for []",
		head:   nil,
		result: false,
	},
}
