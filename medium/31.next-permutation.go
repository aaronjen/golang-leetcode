package main

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutationReverse(nums []int, start, end int) {
	end--
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func nextPermutation(nums []int) {
	n := len(nums)
	k := n - 2
	l := n - 1
	for ; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}
	if k < 0 {
		nextPermutationReverse(nums, 0, n)
	} else {
		for ; l > k; l-- {
			if nums[k] < nums[l] {
				break
			}
		}
		nums[k], nums[l] = nums[l], nums[k]
		nextPermutationReverse(nums, k+1, n)
	}

}

// @lc code=end
