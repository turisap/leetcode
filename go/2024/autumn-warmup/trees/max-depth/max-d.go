package max_depth

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	Node  *TreeNode
	Depth int
}

type Queue struct {
	els []Element
}

func (q *Queue) Len() int {
	return len(q.els)
}

func (q *Queue) Push(element Element) {
	q.els = append(q.els, element)
}

func (q *Queue) Pop() Element {
	el := q.els[len(q.els)-1]
	q.els = q.els[0 : len(q.els)-1]

	return el
}

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepthRecursive(root.Left) + 1
	rd := maxDepthRecursive(root.Right) + 1

	return int(math.Max(float64(ld), float64(rd)))
}

func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{}
	q.Push(Element{root, 1})
	m := 0

	for q.Len() > 0 {
		curr := q.Pop()

		if curr.Depth > m {
			m = curr.Depth
		}

		if curr.Node.Left != nil {
			q.Push(Element{curr.Node.Left, curr.Depth + 1})
		}
		if curr.Node.Right != nil {
			q.Push(Element{curr.Node.Right, curr.Depth + 1})
		}
	}

	return m
}
