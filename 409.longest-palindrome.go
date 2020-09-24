package main

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	table := map[rune]int{}

	for _, r := range s {
		table[r]++
	}

	count := 0
	plusOne := false
	for _, v := range table {
		if v%2 == 0 {
			count += v
		} else {
			count += v - 1
			plusOne = true
		}
	}

	if plusOne {
		count++
	}

	return count
}

// @lc code=end
