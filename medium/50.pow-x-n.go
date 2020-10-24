package main

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	ret := x
	base := 1
	for base*2 < n {
		ret *= ret
		base *= 2
	}
	n -= base

	for i := 0; i < n; i++ {
		ret *= x
	}

	return ret
}

// @lc code=end
