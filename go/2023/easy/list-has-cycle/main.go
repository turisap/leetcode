package main

func lhsTwoPointers(head *ListNode) (res bool) {
	if head == nil {
		return false
	}

	s := head
	f := head.Next

	for s != f {

		if s == nil || f == nil {
			return false
		}

		s = s.Next

		if f.Next != nil {
			f = f.Next.Next
		} else {
			f = nil
		}
	}

	return true
}

func lhsIterative(head *ListNode) bool {
	ht := map[*ListNode]bool{}

	for head != nil {
		if _, ok := ht[head]; !ok {
			ht[head] = true
		} else {
			return true
		}

		head = head.Next
	}

	return false
}
