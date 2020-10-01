package main

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		ind := nums[i]
		if ind < 0 {
			ind *= -1
		}

		if nums[ind-1] > 0 {
			nums[ind-1] *= -1
		}
	}

	ret := []int{}

	for i, v := range nums {
		if v > 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}

// @lc code=end
