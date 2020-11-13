package main

import "container/list"

/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	dis := make([][]int, m)
	queue := list.New()
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				dis[i][j] = 0
				queue.PushBack([]int{i, j})
			} else {
				dis[i][j] = 10001
			}
		}
	}

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for queue.Len() != 0 {
		f := queue.Front()
		cell := f.Value.([]int)
		cellDis := dis[cell[0]][cell[1]]
		for _, d := range dirs {
			r := cell[0] + d[0]
			c := cell[1] + d[1]

			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}

			if dis[r][c] > cellDis+1 {
				dis[r][c] = cellDis + 1
				queue.PushBack([]int{r, c})
			}
		}
		queue.Remove(f)
	}

	return dis
}

// @lc code=end
