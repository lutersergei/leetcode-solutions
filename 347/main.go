package main

import (
	"container/heap"
	"fmt"
)

/*
Top K Frequent Elements
*/
func main() {
	fmt.Println(topKFrequent([]int{-1, -1, 1, 2, 2, 3, 4, 4, 4}, 3))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

func topKFrequent(nums []int, k int) []int {
	l := len(nums)
	res := make([]int, 0, k)
	m := make(map[int]int)
	sl := make([][]int, l)

	for _, num := range nums {
		m[num]++
	}

	for num, count := range m {
		sl[l-count] = append(sl[l-count], num)
	}

	i := 0
	for len(res) < k && i < l {
		res = append(res, sl[i]...)
		i++
	}

	return res
}

func topKFrequent2(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
	}

	q := &priorityQueue{}
	heap.Init(q)
	for val, cnt := range seen {
		heap.Push(q, element{value: val, count: cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int, k, k)
	for i := 1; i <= k; i++ {
		v := heap.Pop(q)
		if s, ok := v.(element); ok {
			ans[k-i] = s.value
		}
	}
	return ans
}

type element struct {
	value int
	count int
}

type priorityQueue []element

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].count < q[j].count }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
