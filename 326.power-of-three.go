package main

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	maxN := 1162261467

	return (n > 0 && (maxN%n == 0))
}

// @lc code=end
