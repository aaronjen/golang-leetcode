package main

/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	lo := 0
	hi := len(arr) - k

	for lo < hi {
		mid := (lo + hi) / 2
		if x-arr[mid] > arr[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return arr[lo : lo+k]
}

// @lc code=end
