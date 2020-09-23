package main

/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// @lc code=end
