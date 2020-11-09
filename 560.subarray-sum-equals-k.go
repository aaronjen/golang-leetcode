package main

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	hash := map[int]int{}
	sum := 0
	cnt := 0
	hash[0] = 1
	for _, n := range nums {
		sum += n
		cnt += hash[sum-k]
		hash[sum]++
	}
	return cnt
}

// @lc code=end
