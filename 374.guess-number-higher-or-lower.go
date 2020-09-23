package main

/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

func guess(num int) int {
	return 0
}

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	max := n
	min := 1

	for {
		v := (max + min) / 2
		ans := guess(v)
		if ans == -1 {
			max = v - 1
		} else if ans == 1 {
			min = v + 1
		} else {
			return v
		}
	}
}

// @lc code=end
