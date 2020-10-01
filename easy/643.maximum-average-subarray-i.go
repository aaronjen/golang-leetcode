package main

/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	cur := 0
	for j := 0; j < k; j++ {
		cur += nums[j]
	}

	max := cur

	for i := 0; i < len(nums)-k; i++ {
		cur += nums[i+k] - nums[i]

		if cur > max {
			max = cur
		}
	}

	return float64(max) / float64(k)
}

// @lc code=end
