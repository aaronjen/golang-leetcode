package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {

	table := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for _, c := range s {
		if cl, ok := table[byte(c)]; ok {
			n := len(stack)
			if len(stack) == 0 || stack[n-1] != cl {
				return false
			}
			stack = stack[:n-1]
		} else {
			stack = append(stack, byte(c))
		}
	}

	return len(stack) == 0
}

// @lc code=end
