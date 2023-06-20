package main

import "testing"

var t = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: nil,
			},
		},
	},
	Right: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: nil,
				},
			},
		},
	},
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = inorderTraversal(t)
	}
}

func BenchmarkAppend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = inorderTraversalOneAppend(t)
	}
}
