package main

import "unicode"

/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	passed := true
	for _, v := range word {
		if !unicode.IsUpper(v) {
			passed = false
			break
		}
	}

	if passed {
		return true
	}

	passed = true
	for _, v := range word {
		if !unicode.IsLower(v) {
			passed = false
			break
		}
	}

	if passed {
		return true
	}

	if !unicode.IsUpper(rune(word[0])) {
		return false
	}
	for i := 1; i < len(word); i++ {
		if !unicode.IsLower(rune(word[i])) {
			return false
		}
	}

	return true
}

// @lc code=end
