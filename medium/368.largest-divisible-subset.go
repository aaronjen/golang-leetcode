package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)

	if n < 1 {
		return []int{}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	sort.Ints(nums)

	dp[0][0] = 1
	dp[0][1] = 0

	gmax := 0
	entry := 0
	for i := 1; i < n; i++ {
		max := 1
		ind := i
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if max < dp[j][0]+1 {
					max = dp[j][0] + 1
					ind = j
				}
			}
		}
		dp[i][0] = max
		dp[i][1] = ind
		if max > gmax {
			gmax = max
			entry = i
		}
	}

	res := []int{}

	for {
		res = append(res, nums[entry])
		if entry == dp[entry][1] {
			break
		}
		entry = dp[entry][1]
	}

	return res
}

// @lc code=end
