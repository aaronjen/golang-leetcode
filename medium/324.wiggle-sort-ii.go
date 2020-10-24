package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */

// @lc code=start
func wiggleSort(nums []int) {
	n := len(nums)
	slice := make([]int, n)
	copy(slice, nums)
	sort.Ints(slice)

	mid := slice[n/2]

	left := 0
	i := 0
	right := n - 1

	for i <= right {
		if nums[newIndex(i, n)] > mid {
			nums[newIndex(i, n)], nums[newIndex(left, n)] = nums[newIndex(left, n)], nums[newIndex(i, n)]
			i++
			left++
		} else if nums[newIndex(i, n)] < mid {
			nums[newIndex(i, n)], nums[newIndex(right, n)] = nums[newIndex(right, n)], nums[newIndex(i, n)]
			right--
		} else {
			i++
		}
	}
}

func newIndex(i, n int) int {
	return (1 + i*2) % (n | 1)
}

// @lc code=end
