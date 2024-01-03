package main

func pllIterative(head *ListNode) bool {
	arr := []*ListNode{}

	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	l, r := 0, len(arr)-1

	for l <= r {
		if arr[l].Val != arr[r].Val {
			return false
		}

		l++
		r--
	}

	return true
}
