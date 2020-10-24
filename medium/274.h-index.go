package main

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)

	for _, c := range citations {
		if c > n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}

	cnt := 0
	for i := n; i >= 0; i-- {
		cnt += buckets[i]
		if cnt >= i {
			return i
		}

	}

	return 0
}

// @lc code=end
