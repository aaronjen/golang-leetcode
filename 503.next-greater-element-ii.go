package main

/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	stack := []int{}

	for i := 0; i < n*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}

	}

	return res
}

// @lc code=end
