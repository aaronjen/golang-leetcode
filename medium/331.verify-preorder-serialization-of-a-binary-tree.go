package main

import "strings"

/*
 * @lc app=leetcode id=331 lang=golang
 *
 * [331] Verify Preorder Serialization of a Binary Tree
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	arr := strings.Split(preorder, ",")

	diff := 1

	for _, s := range arr {
		diff--
		if diff < 0 {
			return false
		}

		if s != "#" {
			diff += 2
		}
	}

	return diff == 0
}

// @lc code=end
