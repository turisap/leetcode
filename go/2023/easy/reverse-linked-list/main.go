package main

func rllRecursive(head *ListNode) *ListNode {
	current := head

	if head == nil || head.Next == nil {
		return head
	}

	newHead := rllRecursive(current.Next)

	current.Next.Next = current
	current.Next = nil

	return newHead
}

func rllIterative(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var revHead *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}

	return revHead
}
