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

/*
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkAppend
BenchmarkAppend-8        1274215               897.8 ns/op           784 B/op         29 allocs/op
BenchmarkAppend2
BenchmarkAppend2-8       1291267               911.4 ns/op           784 B/op         29 allocs/op
*/

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
