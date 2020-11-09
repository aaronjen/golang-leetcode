package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	if S > 1000 {
		return 0
	}

	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2001)
	}

	dp[0][1000] = 1
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := 0; j < 2001; j++ {
			prevCnt := dp[i][j]
			if prevCnt != 0 {
				dp[i+1][j+v] += prevCnt
				dp[i+1][j-v] += prevCnt
			}
		}
	}

	return dp[n][1000+S]
}

// @lc code=end
