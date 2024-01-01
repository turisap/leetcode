package main

func lliHashTable(h1, h2 *ListNode) *ListNode {
	ht := map[*ListNode]bool{}

	for h1 != nil {
		ht[h1] = true
		h1 = h1.Next
	}

	for h2 != nil {
		if _, ok := ht[h2]; ok {
			return h2
		}
		h2 = h2.Next
	}

	return nil
}

func getLength(h *ListNode) (l int) {
	c := h
	for c != nil {
		l++
		c = c.Next
	}

	return
}

func lliTwoPointersImproved(h1, h2 *ListNode) *ListNode {
	hA := h1
	hB := h2
	l1 := getLength(h1)
	l2 := getLength(h2)

	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			hA = hA.Next
		}
	} else if l2 > l1 {
		for i := 0; i < l2-l1; i++ {
			hB = hB.Next
		}
	}

	for hA != nil && hB != nil {
		if hA == hB {
			return hB
		}

		hA = hA.Next
		hB = hB.Next
	}

	return nil
}

func lliTwoPointersPlain(h1, h2 *ListNode) *ListNode {
	l1 := getLength(h1)
	l2 := getLength(h2)

	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			h1 = h1.Next
		}
	} else if l2 > l1 {
		for i := 0; i < l2-l1; i++ {
			h2 = h2.Next
		}
	}

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h2
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	return nil
}
