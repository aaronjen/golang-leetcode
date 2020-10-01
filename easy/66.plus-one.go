package main

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	ret := []int{}

	n := len(digits)

	payload := true
	for i := 0; i < n; i++ {
		v := digits[n-i-1]
		if payload {
			v = v + 1
			if v > 9 {
				v = 0
				payload = true
			} else {
				payload = false
			}
		}
		ret = append(ret, v)
	}

	if payload {
		ret = append(ret, 1)
	}

	l := len(ret)

	for i := 0; i < l/2; i++ {
		ret[i], ret[l-i-1] = ret[l-i-1], ret[i]
	}

	return ret
}

// @lc code=end
