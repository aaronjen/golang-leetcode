package main

import "sort"

/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 */

// @lc code=start

func findLongestWord(s string, d []string) string {
	sort.Strings(d)
	res := ""
	for _, it := range d {
		if len(it) <= len(res) {
			continue
		}

		i := 0
		j := 0
		for j < len(s) {
			if it[i] == s[j] {
				i++
				j++
				if i >= len(it) {
					res = it
					break
				}
			} else {
				j++
			}
		}
	}

	return res
}

// @lc code=end
