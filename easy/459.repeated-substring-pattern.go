package main

import (
	"bytes"
)

/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sub := s[0:i]
			times := n / i

			var buffer bytes.Buffer

			for i := 0; i < times; i++ {
				buffer.WriteString(sub)
			}

			if s == buffer.String() {
				return true
			}
		}
	}

	return false
}

// @lc code=end
