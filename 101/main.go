package main

import "fmt"

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	fmt.Println(isSymmetric(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var f func(n1, n2 *TreeNode) bool

	f = func(n1, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil && n2 != nil || n1 != nil && n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		return f(n1.Left, n2.Right) && f(n1.Right, n2.Left)
	}

	return f(root.Left, root.Right)
}
