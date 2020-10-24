package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslandsMarking(grid [][]byte, x, y int) {
	ly := len(grid)
	lx := len(grid[0])
	if x < 0 || y < 0 || x >= lx || y >= ly || grid[y][x] != '1' {
		return
	}
	grid[y][x] = '0'
	numIslandsMarking(grid, x-1, y)
	numIslandsMarking(grid, x+1, y)
	numIslandsMarking(grid, x, y-1)
	numIslandsMarking(grid, x, y+1)

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	if len(grid[0]) == 0 {
		return 0
	}

	ly := len(grid)
	lx := len(grid[0])
	cnt := 0
	for y := 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if grid[y][x] == '1' {
				cnt++
				numIslandsMarking(grid, x, y)
			}
		}
	}

	return cnt
}

// @lc code=end
