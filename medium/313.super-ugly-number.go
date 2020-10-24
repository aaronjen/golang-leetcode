package main

import (
	"math"
)

/*
 * @lc app=leetcode id=313 lang=golang
 *
 * [313] Super Ugly Number
 */

// @lc code=start
func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	idx := make([]int, len(primes))

	ugly[0] = 1
	for i := 1; i < n; i++ {
		ugly[i] = math.MaxInt32
		for j := 0; j < len(primes); j++ {
			v := ugly[idx[j]] * primes[j]
			if v < ugly[i] {
				ugly[i] = v
			}
		}
		for j := 0; j < len(primes); j++ {
			for ugly[idx[j]]*primes[j] <= ugly[i] {
				idx[j]++
			}
		}
	}

	return ugly[n-1]
}

// @lc code=end
