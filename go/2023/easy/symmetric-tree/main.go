package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sytRecursive(root *TreeNode) bool {
	var inner func(l, r *TreeNode) bool

	inner = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if (l == nil && r != nil) || (l != nil && r == nil) {
			return false
		}

		return l.Val == r.Val && inner(l.Left, r.Right) && inner(l.Right, r.Left)
	}

	return inner(root.Left, root.Right)
}

type Queue struct {
	items []*TreeNode
}

func (q *Queue) DequeueTwo() (r, l *TreeNode) {
	if len(q.items) >= 2 {
		l = q.items[0]
		r = q.items[1]
		q.items = q.items[2:]

		return r, l
	}

	return nil, nil
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Enqueue(node []*TreeNode) {
	q.items = append(q.items, node...)
}

func sytIterative(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	var q Queue
	q.Enqueue([]*TreeNode{root.Left, root.Right})

	for q.Len() > 0 {
		l, r := q.DequeueTwo()

		if l == nil && r != nil {
			return false
		}

		if l != nil && r == nil {
			return false
		}

		if l != nil && r != nil {
			if l.Val != r.Val {
				return false
			}
			q.Enqueue([]*TreeNode{l.Left, r.Right, l.Right, r.Left})
		}
	}

	return true
}
