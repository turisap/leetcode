package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var c1 = &ListNode{Val: 2, Next: &ListNode{Val: 4}}
var c2 = &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 14}}}

var testCases = []struct {
	name   string
	head1  *ListNode
	head2  *ListNode
	result *ListNode
}{
	{
		"false for listA = [2,6,4], listB = [1,5]",
		&ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 5}},
		nil,
	},
	{
		"true for listA = [1,9,1,2,4], listB = [3,2,4]",
		&ListNode{Val: 5, Next: &ListNode{Val: 9, Next: &ListNode{Val: 2, Next: c1}}},
		&ListNode{Val: 3, Next: c1},
		c1,
	},
	{
		"true for listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]",
		&ListNode{Val: 19, Next: &ListNode{Val: 1, Next: c2}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: c2}}},
		c2,
	},
}
