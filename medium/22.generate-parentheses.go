package main

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesisSolve(cur string, n int, open int, ret *[]string) {
	if n == 0 && open == 0 {
		*ret = append(*ret, cur)
	}

	if n > 0 {
		s := cur + "("
		generateParenthesisSolve(s, n-1, open+1, ret)
	}
	if open > 0 {
		s := cur + ")"
		generateParenthesisSolve(s, n, open-1, ret)
	}
}

func generateParenthesis(n int) []string {
	ret := []string{}

	generateParenthesisSolve("", n, 0, &ret)

	return ret
}

// @lc code=end
