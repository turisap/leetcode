package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var testCases = []struct {
	name   string
	head   *ListNode
	result *ListNode
}{
	{
		"[1,2,3,4,5]",
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
	},
}
