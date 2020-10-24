package main

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := [][]int{
		intervals[0],
	}

	for i := 1; i < len(intervals); i++ {
		n := len(ret)
		last := ret[n-1]
		cur := intervals[i]

		if cur[0] <= last[1] {
			e := last[1]
			if e < cur[1] {
				e = cur[1]
			}

			ret[n-1] = []int{last[0], e}
		} else {
			ret = append(ret, cur)
		}
	}

	return ret
}

// @lc code=end
