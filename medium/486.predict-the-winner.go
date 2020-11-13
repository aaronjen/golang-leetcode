package main

/*
 * @lc app=leetcode id=486 lang=golang
 *
 * [486] Predict the Winner
 */

// @lc code=start

func solve(nums []int, i, j int, dp [][]int) int {
	if i > j {
		return 0
	}

	if i == j {
		return nums[i]
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	a := solve(nums, i+1, j-1, dp)
	b := solve(nums, i+2, j, dp)
	v1 := nums[i] + a
	if a > b {
		v1 = nums[i] + b
	}

	c := solve(nums, i+1, j-1, dp)
	d := solve(nums, i, j-2, dp)
	v2 := nums[j] + c
	if c > d {
		v2 = nums[j] + d
	}

	max := v1
	if v1 < v2 {
		max = v2
	}

	dp[i][j] = max

	return max
}

// PredictTheWinner func
func PredictTheWinner(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	solve(nums, 0, n-1, dp)

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	return dp[0][n-1] >= sum-dp[0][n-1]
}

// @lc code=end
