package main

/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	n := len(num)

	if k == n {
		return "0"
	}

	stack := []byte{}

	for i := 0; i < n; i++ {
		c := num[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, c)
	}

	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}

	return string(stack)
}

// @lc code=end
