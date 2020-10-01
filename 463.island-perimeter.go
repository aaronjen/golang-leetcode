package main

/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	ly := len(grid)
	lx := len(grid[0])

	ret := 0

	for i := 0; i < ly; i++ {
		for j := 0; j < lx; j++ {
			if grid[i][j] == 1 {
				cnt := 0
				if i == 0 || grid[i-1][j] == 0 {
					cnt++
				}

				if i == ly-1 || grid[i+1][j] == 0 {
					cnt++
				}

				if j == 0 || grid[i][j-1] == 0 {
					cnt++
				}

				if j == lx-1 || grid[i][j+1] == 0 {
					cnt++
				}

				ret += cnt
			}
		}
	}

	return ret
}

// @lc code=end
