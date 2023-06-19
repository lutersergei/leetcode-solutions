package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("Val: %v, Left: %v, Right: %v", t.Val, t.Left.Val, t.Right.Val)
}

func main() {
	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	t = invertTree(t)

	fmt.Printf("%v", t)
	fmt.Printf("%v", t.Left)
	fmt.Printf("%v", t.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
