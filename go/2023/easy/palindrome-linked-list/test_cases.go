package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var testCases = []struct {
	name   string
	head   *ListNode
	result bool
}{
	{"true for 1-2-2-1", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}, true},
	{"false for 1-2-2-5", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}, false},
	{"true for 1", &ListNode{Val: 1}, true},
}
