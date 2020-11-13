package main

/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		sum := 0
		for _, n := range nums {
			if n <= i {
				sum += dp[i-n]
			}
		}
		dp[i] = sum
	}

	return dp[target]
}

// @lc code=end
