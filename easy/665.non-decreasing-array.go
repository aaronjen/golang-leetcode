package main

/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	cnt := 0

	for i := 1; i < len(nums) && cnt < 2; i++ {
		if nums[i] < nums[i-1] {
			cnt++
			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}

		}
	}

	return cnt < 2
}

// @lc code=end
