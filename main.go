package main

import (
	"fmt"
	"math"
)

func main() {
	//img := [][]byte{{'0', '1', '0'}, {'1', '1', '1'}}
	//img := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	//fmt.Println(numIslands(img))
	//fmt.Println(twoSum([]int{-3, 12, 0, -1}, 11))
	//fmt.Println(maxSubArray([]int{-2, -1, -4}))
	//fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Println(fib(5))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(minCostClimbingStairs([]int{6, 5, 4, 3, 2, 1}))
}

func minCostClimbingStairs(cost []int) int {
	var sum int

	l := len(cost)
	m := make(map[int]int, l)
	m[0] = cost[0]
	m[1] = cost[1]

	sum = getCost(cost, l, m)

	return sum
}

func getCost(cost []int, ind int, m map[int]int) int {
	var c int
	prev := getCost(cost, ind-1, m)
	prev2 := getCost(cost, ind-2, m)
	if ind == len(cost) {
		c = min(prev, prev2)
	} else {
		c = min(prev, prev2) + cost[ind]
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMin(cost []int, a, b int, m map[int]int) int {
	aCost, ok := m[a]
	if !ok {
		aCost = cost[a] + getMin(cost, a-1, a-2, m)
		m[a] = aCost
	}

	bCost, ok := m[a]
	if !ok {
		bCost = cost[b] + getMin(cost, b-1, b-2, m)
		m[b] = bCost
	}

	if aCost < bCost {
		return aCost
	}
	return bCost
}

//func fib(n int) int {
//	if n <= 1 {
//		return n
//	}
//
//	return fib(n-1) + fib(n-2)
//}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	f, s := 0, 1
	for i := 0; i < n-1; i++ {
		f, s = s, f+s
	}
	return s
}

//	func numIslands(grid [][]byte) int {
//		rows, columns := len(grid), len(grid[0])
//		m := make(map[[2]int]struct{})
//		var count int
//		var traversIsl func(int, int) bool
//
//		traversIsl = func(i, i2 int) bool {
//			if _, ok := m[[2]int{i, i2}]; ok {
//				return false
//			}
//			if grid[i][i2] == '0' {
//				return false
//			}
//			m[[2]int{i, i2}] = struct{}{}
//			if i < rows-1 {
//				traversIsl(i+1, i2)
//			}
//			if i2 < columns-1 {
//				traversIsl(i, i2+1)
//			}
//			if i > 0 {
//				traversIsl(i-1, i2)
//			}
//			if i2 > 0 {
//				traversIsl(i, i2-1)
//			}
//			return true
//		}
//
//		for r, row := range grid {
//			for c := range row {
//				if traversIsl(r, c) {
//					count++
//				}
//			}
//		}
//
//		return count
//	}
func numIslands(grid [][]byte) int {
	rows, columns := len(grid), len(grid[0])
	var count int
	var demark func(int, int)

	demark = func(i, i2 int) {
		if i < 0 || i2 < 0 || i >= rows || i2 >= columns {
			return
		}
		if grid[i][i2] == '0' {
			return
		}

		grid[i][i2] = '0'

		demark(i+1, i2)
		demark(i, i2+1)
		demark(i-1, i2)
		demark(i, i2-1)
	}

	for r, row := range grid {
		for c := range row {
			if grid[r][c] == '1' {
				count++
				demark(r, c)
			}
		}
	}

	return count
}

func markAsIsland(i, i2 int, m map[[2]int]struct{}, grid [][]byte) bool {
	rows, columns := len(grid), len(grid[0])

	if _, ok := m[[2]int{i, i2}]; ok {
		return false
	}
	if grid[i][i2] == 0 {
		return false
	}

	m[[2]int{i, i2}] = struct{}{}

	if i < rows-1 {
		markAsIsland(i+1, i2, m, grid)
	}
	if i2 < columns-1 {
		markAsIsland(i, i2+1, m, grid)
	}

	return true
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, item := range nums {
		if ind, ok := m[target-item]; ok {
			return []int{ind, i}
		}
		m[item] = i
	}
	return []int{}
}

func maxSubArray(nums []int) int {
	maxSum, tmpSum := nums[0], 0

	for _, item := range nums {
		tmpSum += item
		if tmpSum > maxSum {
			maxSum = tmpSum
		}
		if tmpSum < 0 {
			tmpSum = 0
		}
	}
	return maxSum
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	oldColor := image[sr][sc]
	image[sr][sc] = color

	if checkExistAndVal(image, sr, sc, sr+1, sc, oldColor) {
		image = floodFill(image, sr+1, sc, color)
	}
	if checkExistAndVal(image, sr, sc, sr, sc+1, oldColor) {
		image = floodFill(image, sr, sc+1, color)
	}
	if checkExistAndVal(image, sr, sc, sr-1, sc, oldColor) {
		image = floodFill(image, sr-1, sc, color)
	}
	if checkExistAndVal(image, sr, sc, sr, sc-1, oldColor) {
		image = floodFill(image, sr, sc-1, color)
	}
	return image
}

func checkExistAndVal(image [][]int, sr, sc, tr, tc, oldColor int) bool {
	if tr < 0 || tc < 0 {
		return false
	}
	if len(image)-1 >= tr && len(image[tr])-1 >= tc {
		if oldColor == image[tr][tc] {
			return true
		}
	}
	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	return checkBST(root, -math.MaxInt64, math.MaxInt64)
}

func checkBST(n *TreeNode, min, max int) bool {
	if n == nil {
		return true
	}

	if n.Val > min && n.Val < max {
		return checkBST(n.Left, min, n.Val) && checkBST(n.Right, n.Val, max)
	}

	return false
}

func search(nums []int, target int) int {
	var s func(start, end int, nums []int, target int) int

	s = func(start, end int, nums []int, target int) int {
		if end == start {
			if nums[end] != target {
				return -1
			}
			return end
		}

		if len(nums) == 0 {
			return -1
		}

		middleIdx := start + ((end - start) / 2)

		if nums[middleIdx] == target {
			return middleIdx
		} else if nums[middleIdx] < target {
			return s(middleIdx+1, end, nums, target)
		} else {
			return s(start, middleIdx, nums, target)
		}
	}

	index := s(0, len(nums)-1, nums, target)

	return index
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var level []int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		level, queue = travers(queue)
		result = append(result, level)
	}

	return result
}

func travers(queue []*TreeNode) ([]int, []*TreeNode) {
	var level []int
	var q []*TreeNode
	for i := 0; i < len(queue); i++ {
		if queue[i] != nil {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				q = append(q, queue[i].Left)
			}
			if queue[i].Right != nil {
				q = append(q, queue[i].Right)
			}
		}
	}

	return level, q
}

func longestPalindrome(s string) int {
	hash := make(map[int32]int)
	var middle bool
	var res int
	for _, i2 := range s {
		hash[i2]++
	}

	for _, count := range hash {
		if count%2 == 0 {
			res += count
		} else if middle {
			res += count - 1
		} else {
			res += count
			middle = true
		}
	}
	return res
}

func split(nums []int, target int) {

}

func getChilds(n *TreeNode) []int {
	return []int{n.Left.Val, n.Right.Val}
}

func maxProfit(prices []int) int {
	var result, tmp int
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	for _, price := range prices {
		tmp = price - min
		if price > min && tmp > result {
			result = tmp
		}
		if price < min {
			min = price
		}
	}
	return result
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var slow, fast *ListNode = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		} else {
			m[head] = struct{}{}
		}
		head = head.Next
	}

	return nil
}

func removeElements(head *ListNode, val int) *ListNode {
	p := head

	var prev *ListNode

	for head != nil {
		if head.Val == val {
			if prev == nil {
				p = head.Next
			} else {
				prev.Next = head.Next
			}
		} else {
			prev = head
		}

		head = head.Next
	}

	return p
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return p
}

func largestAltitude(gain []int) int {
	var max, sum int
	for _, val := range gain {
		sum += val
		if sum > max {
			max = sum
		}
	}
	return max
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[int]*ListNode)
	var count, middle int
	for head != nil {
		count += 1
		m[count] = head
		head = head.Next
	}

	middle = (count / 2) + 1
	return m[middle]
}

func reverseList(head *ListNode) *ListNode {
	var resHead, tmp *ListNode

	for head != nil {
		tmp = head.Next
		head.Next = resHead
		resHead = head
		head = tmp
	}

	return resHead
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}

	result := new(ListNode)

	tail := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	for list1 != nil {
		tail.Next = list1
		tail = tail.Next
		list1 = list1.Next
	}

	for list2 != nil {
		tail.Next = list2
		tail = tail.Next
		list2 = list2.Next
	}

	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//func isSubsequence(s string, t string) bool {
//	j := 0
//	var res []rune
//	for _, char := range s {
//		for i := j; i < len(t); i++ {
//			if char == rune(t[i]) {
//				res = append(res, char)
//				j = i + 1
//				break
//			}
//		}
//	}
//
//	return len(s) == len(res)
//}

//func isIsomorphic(s string, t string) bool {
//	hashTableS := make(map[rune]rune)
//	hashTableT := make(map[rune]rune)
//
//	sRune := []rune(s)
//	tRune := []rune(t)
//
//	for i := 0; i < len(sRune); i++ {
//		if r, ok := hashTableT[tRune[i]]; ok {
//			if r != sRune[i] {
//				return false
//			}
//		}
//		if r, ok := hashTableS[sRune[i]]; ok {
//			if r != tRune[i] {
//				return false
//			}
//		}
//		hashTableT[tRune[i]] = sRune[i]
//		hashTableS[sRune[i]] = tRune[i]
//	}
//
//	return true
//}

//func sortedSquares(nums []int) []int {
//	right := len(nums) - 1
//	left := 0
//	var result []int
//	var lSQ, rSQ int
//	for left <= right {
//		lSQ = nums[left] * nums[left]
//		rSQ = nums[right] * nums[right]
//		if lSQ > rSQ {
//			result = append([]int{lSQ}, result...)
//			left++
//		} else {
//			result = append([]int{rSQ}, result...)
//			right--
//		}
//		//n--
//	}
//
//	return result
//}

//func isIsomorphic(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	var charMapS, charMapT [1280]uint8
//
//	for i := range s {
//		charS, charT := s[i], t[i]
//		if (charMapS[charS] != 0 && charMapS[charS] != charT) ||
//			(charMapT[charT] != 0 && charMapT[charT] != charS) {
//			return false
//		} else {
//			charMapS[charS], charMapT[charT] = charT, charS
//		}
//	}
//
//	return true
//}

//func pivotIndex(nums []int) int {
//	sumLeft := 0
//	sumRight := 0
//
//	for _, num := range nums {
//		sumRight += num
//	}
//
//	for i := range nums {
//		sumRight -= nums[i]
//
//		if sumLeft == sumRight {
//			return i
//		}
//
//	}
//
//	return -1
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
