package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	ret := [][]int{}
	s := newInterval[0]
	e := newInterval[1]
	inserted := false
	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]

		if inserted {
			ret = append(ret, cur)
		} else if cur[1] < s {
			ret = append(ret, cur)
		} else if cur[0] > e {
			ret = append(ret, []int{s, e}, cur)
			inserted = true
		} else {
			if cur[0] < s {
				s = cur[0]
			}

			if cur[1] > e {
				e = cur[1]
			}
		}
	}

	if inserted == false {
		ret = append(ret, []int{s, e})
	}

	return ret
}

// @lc code=end
