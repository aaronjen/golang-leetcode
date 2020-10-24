package main

/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2
		if mid != len(nums)-1 && nums[mid] < nums[mid+1] {
			left = mid + 1
			continue
		}

		if mid != 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
			continue
		}
		return mid
	}
}

// @lc code=end
