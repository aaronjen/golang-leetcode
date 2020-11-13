package main

/*
 * @lc app=leetcode id=481 lang=golang
 *
 * [481] Magical String
 */

// @lc code=start
func magicalString(n int) int {
	s := "122"
	ind := 2
	for len(s) < n {
		r := int(s[ind] - '0')
		c := s[len(s)-1] ^ 3
		for j := 0; j < r; j++ {
			s += string(c)
		}
		ind++
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			cnt++
		}
	}

	return cnt
}

// @lc code=end
