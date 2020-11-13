package main

import (
	"math"
)

/*
 * @lc app=leetcode id=375 lang=golang
 *
 * [375] Guess Number Higher or Lower II
 */

// @lc code=start
func getMoneyAmount(n int) int {
	table := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]int, n+1)
	}

	return getMoneyAmountDP(table, 1, n)
}

func getMoneyAmountDP(table [][]int, s int, e int) int {
	if s >= e {
		return 0
	}
	if table[s][e] != 0 {
		return table[s][e]
	}

	res := math.MaxInt32
	for i := s; i <= e; i++ {
		a := getMoneyAmountDP(table, s, i-1)
		b := getMoneyAmountDP(table, i+1, e)
		var tmp int
		if a > b {
			tmp = i + a
		} else {
			tmp = i + b
		}
		if tmp < res {
			res = tmp
		}
	}

	table[s][e] = res
	return res
}

// @lc code=end
