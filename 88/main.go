package main

import "fmt"

// Merge Sorted Array
func main() {
	a := []int{1, 2, 3, 8, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 4, b, 3)
	fmt.Println(a)

	a = []int{0}
	b = []int{1}
	merge(a, 0, b, 1)
	fmt.Println(a)

	a = []int{1}
	b = []int{}
	merge(a, 1, b, 0)
	fmt.Println(a)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	lastP := m + n - 1
	m -= 1
	n -= 1

	for n >= 0 && m >= 0 {
		if nums1[m] > nums2[n] {
			nums1[lastP] = nums1[m]
			lastP--
			m--
			continue
		}

		nums1[lastP] = nums2[n]
		lastP--
		n--
	}

	for n >= 0 {
		nums1[lastP] = nums2[n]
		lastP -= 1
		n -= 1
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1[:], nums2[:])
		return
	}

	var i, j int

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			i++
			continue
		}

		copy(nums1[i+1:m+1], nums1[i:m])
		nums1[i] = nums2[j]
		i++
		j++
		m++
	}

	if m >= i {
		copy(nums1[m:], nums2[j:])
	}
}
