package main

/*
 * @lc app=leetcode id=421 lang=golang
 *
 * [421] Maximum XOR of Two Numbers in an Array
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	max := 0
	mask := 0

	for i := 31; i >= 0; i-- {
		mask = mask | (1 << i)

		set := map[int]bool{}
		for _, num := range nums {
			set[num&mask] = true
		}

		tmp := max | (1 << i)

		for prefix := range set {
			if ok := set[tmp^prefix]; ok {
				max = tmp
				break
			}
		}
	}

	return max
}

// @lc code=end
