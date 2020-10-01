package main

/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	if k <= 1 {
		return s
	}

	ret := []byte{}

	ls := len(s)
	pad := 0
	for pad+2*k < ls {
		for i := 0; i < k; i++ {
			ret = append(ret, s[pad+k-i-1])
		}

		for i := 0; i < k; i++ {
			ret = append(ret, s[pad+k+i])
		}

		pad += 2 * k
	}

	left := ls - pad
	if left > k {
		left = k
	}

	for i := 0; i < left; i++ {
		ret = append(ret, s[pad+left-i-1])
	}

	for i := pad + left; i < ls; i++ {
		ret = append(ret, s[i])
	}

	return string(ret)
}

// @lc code=end
