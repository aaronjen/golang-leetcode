package main

/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				a := dp[i-1][j]
				b := dp[i][j-1]

				if a > b {
					dp[i][j] = a
				} else {
					dp[i][j] = b
				}
			}
		}
	}

	v := dp[m][n]

	return m + n - 2*v
}

// @lc code=end
