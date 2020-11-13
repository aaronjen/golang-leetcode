package main

import (
	"math"
)

/*
 * @lc app=leetcode id=659 lang=golang
 *
 * [659] Split Array into Consecutive Subsequences
 */

// @lc code=start
func isPossible(nums []int) bool {
	pre := math.MinInt16

	p1, p2, p3 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		cnt := 1
		for i+1 < len(nums) && cur == nums[i+1] {
			cnt++
			i++
		}

		if pre+1 != cur {
			if p1 != 0 || p2 != 0 {
				return false
			}
			p1 = cnt
			p2 = 0
			p3 = 0
		} else {
			if cnt < p1+p2 {
				return false
			}

			cntLeft := cnt - p1 - p2
			if cntLeft > p3 {
				need := cntLeft - p3
				p3 = p2 + p3
				p2 = p1
				p1 = need

			} else {
				p3 = p2 + cntLeft
				p2 = p1
				p1 = 0
			}
		}
		pre = cur
	}

	return p1 == 0 && p2 == 0
}

// @lc code=end
