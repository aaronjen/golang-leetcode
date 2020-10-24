package main

import "strconv"

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	start := nums[0]
	end := nums[0]
	res := []string{}

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if v == end+1 {
			end = v
		} else {
			if start == end {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}

			start = v
			end = v
		}
	}

	if start == end {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}

	return res
}

// @lc code=end
