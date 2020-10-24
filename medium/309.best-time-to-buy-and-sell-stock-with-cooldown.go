package main

import (
	"math"
)

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfitMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	s0 := make([]int, n) // rest
	s1 := make([]int, n) // buy
	s2 := make([]int, n) // sell

	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = math.MinInt32

	for i := 1; i < n; i++ {
		s0[i] = maxProfitMax(s0[i-1], s2[i-1])
		s1[i] = maxProfitMax(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	if s0[n-1] >= s1[n-1] && s0[n-1] >= s2[n-1] {
		return s0[n-1]
	}

	if s1[n-1] >= s0[n-1] && s1[n-1] >= s2[n-1] {
		return s1[n-1]
	}

	return s2[n-1]
}

// @lc code=end
