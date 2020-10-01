package main

/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for _, d := range nums {
		if d == 1 {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}

	if cur > max {
		max = cur
	}

	return max
}

// @lc code=end
