package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MTSReg(list1 *ListNode, list2 *ListNode) *ListNode {
	var res = &ListNode{}
	var cur = res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return res.Next
}

func MTSL(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MTSL(list1.Next, list2)
		return list1
	}

	list2.Next = MTSL(list2.Next, list1)
	return list2
}

// c1 -> next
// c2 -> next
