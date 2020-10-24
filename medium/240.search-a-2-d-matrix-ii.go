package main

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	ly := len(matrix)

	if ly == 0 {
		return false
	}
	lx := len(matrix[0])
	if lx == 0 {
		return false
	}

	col := lx - 1
	row := 0
	for col >= 0 && row < ly {
		v := matrix[row][col]
		if v == target {
			return true
		} else if v > target {
			col--
		} else {
			row++
		}
	}

	return false
}

// @lc code=end
