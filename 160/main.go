package main

import (
	"fmt"
)

func main() {
	intList := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  25,
			Next: nil,
		},
	}

	headA := &ListNode{
		Val:  2,
		Next: intList,
	}

	headB := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 43,
			Next: &ListNode{
				Val:  33,
				Next: intList,
			},
		},
	}

	fmt.Println(getIntersectionNode(headA, headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
