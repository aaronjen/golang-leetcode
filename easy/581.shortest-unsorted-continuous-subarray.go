package main

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	end := 0
	beg := len(nums) - 1
	n := len(nums)

	max := nums[0]
	min := nums[n-1]
	for i := 1; i < n; i++ {
		if nums[i] < max {
			end = i
		}
		if nums[i] > max {
			max = nums[i]
		}

		if nums[n-i-1] > min {
			beg = n - i - 1
		}
		if nums[n-i-1] < min {
			min = nums[n-i-1]
		}
	}

	if end < beg {
		return 0
	}

	return end - beg + 1
}

// @lc code=end
