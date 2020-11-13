package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 */

// @lc code=start
func repeatedStringMatch(a string, b string) int {
	n := len(b) / len(a)
	if len(b)%len(a) != 0 {
		n++
	}

	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteString(a)
	}
	if strings.Index(sb.String(), b) != -1 {
		return n
	}

	sb.WriteString(a)
	if strings.Index(sb.String(), b) != -1 {
		return n + 1
	}

	return -1
}

// @lc code=end
