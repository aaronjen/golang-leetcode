package main

/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		tmpDP := make([][]int, m+1)
		for i := 0; i < m+1; i++ {
			tmpDP[i] = make([]int, n+1)
			copy(tmpDP[i], dp[i])
		}
		s := strs[i]
		zs := 0
		os := 0
		for _, b := range s {
			if b == '0' {
				zs++
			} else if b == '1' {
				os++
			}
		}
		for j := zs; j <= m; j++ {
			for k := os; k <= n; k++ {
				cur := dp[j][k]
				tmp := dp[j-zs][k-os] + 1
				if tmp > cur {
					tmpDP[j][k]++
				}
			}
		}
		dp = tmpDP
	}

	return dp[m][n]
}

// @lc code=end
