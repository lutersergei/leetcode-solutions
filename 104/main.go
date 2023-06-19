package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Max binary tree depth
func main() {
	t := &TreeNode{}

	fmt.Println(maxDepth(t))
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
