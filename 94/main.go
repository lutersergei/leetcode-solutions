package main

import "fmt"

func main() {
	t := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: nil,
		},
	}

	fmt.Println(inorderTraversal(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}

func inorderTraversalOneAppend(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(append([]int{}, inorderTraversal(root.Left)...), root.Val), inorderTraversal(root.Right)...)

}

//func traverse(node *TreeNode, res []int) {
//
//}
