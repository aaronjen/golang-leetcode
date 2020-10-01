package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	n := len(nums)
	type pair struct {
		a int
		b int
	}

	order := make([]pair, n)

	for ind, v := range nums {
		order[ind] = pair{
			ind,
			v,
		}
	}

	sort.Slice(order, func(i, j int) bool {
		return order[i].b > order[j].b
	})

	ret := make([]string, n)

	for ind, o := range order {
		rank := strconv.Itoa(ind + 1)
		if ind == 0 {
			rank = "Gold Medal"
		} else if ind == 1 {
			rank = "Silver Medal"
		} else if ind == 2 {
			rank = "Bronze Medal"
		}

		ret[o.a] = rank
	}

	return ret
}

// @lc code=end
