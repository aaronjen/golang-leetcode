package main

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	si := 0
	for _, r := range s {
		si += int(r)
	}

	ti := 0
	for _, r := range t {
		ti += int(r)
	}

	return byte(ti - si)
}

// @lc code=end
