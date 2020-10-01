package main

/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */

// @lc code=start
func judgeCircle(moves string) bool {
	x := 0
	y := 0

	for _, m := range moves {
		switch m {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	return x == 0 && y == 0
}

// @lc code=end
