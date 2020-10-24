package main

/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */

// @lc code=start
func maxProduct(words []string) int {
	l := len(words)

	if l == 0 {
		return 0
	}

	values := make([]int, len(words))

	for i, w := range words {
		v := 0
		for _, c := range w {
			v |= 1 << int(c-'a')
		}

		values[i] = v
	}

	max := 0

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if values[i]&values[j] != 0 {
				continue
			}

			v := len(words[i]) * len(words[j])
			if v > max {
				max = v
			}
		}
	}

	return max
}

// @lc code=end
