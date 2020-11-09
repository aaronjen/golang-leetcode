package main

/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */

// @lc code=start
func integerReplacementSolve(dp map[int]int, n int) int {
	if n == 1 {
		return 0
	}

	if v, ok := dp[n]; ok {
		return v
	}

	if n%2 == 0 {
		v := integerReplacementSolve(dp, n/2) + 1
		dp[n] = v

		return v
	}
	v1 := integerReplacementSolve(dp, n+1) + 1
	v2 := integerReplacementSolve(dp, n-1) + 1
	v := v1
	if v1 > v2 {
		v = v2
	}

	dp[n] = v
	return v
}

func integerReplacement(n int) int {
	dp := map[int]int{}

	return integerReplacementSolve(dp, n)
}

// @lc code=end
