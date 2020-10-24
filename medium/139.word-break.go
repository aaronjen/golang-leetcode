package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreakSolve(s string, wordDict []string, dp []bool) {
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			sub := s[j:i]

			if !dp[j] {
				continue
			}
			for _, w := range wordDict {
				if sub == w {
					dp[i] = true
					break
				}
			}
		}
	}
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordBreakSolve(s, wordDict, dp)

	return dp[len(s)]
}

// @lc code=end
