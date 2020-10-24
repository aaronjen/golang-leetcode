package main

import "math"

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	dp := []int{0}

	for len(dp) <= n {
		m := len(dp)
		min := math.MaxInt16

		for i := 1; i*i <= m; i++ {
			val := dp[m-i*i] + 1
			if min > val {
				min = val
			}
		}
		dp = append(dp, min)
	}

	return dp[n]
}

// @lc code=end
