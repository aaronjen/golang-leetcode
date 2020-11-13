package main

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func buildTable(matrix [][]int, table [][]bool, x, y int) {
	if table[y][x] == true {
		return
	}

	table[y][x] = true
	ly := len(matrix)
	lx := len(matrix[0])

	curH := matrix[y][x]

	if x != 0 && matrix[y][x-1] >= curH {
		buildTable(matrix, table, x-1, y)
	}

	if x != lx-1 && matrix[y][x+1] >= curH {
		buildTable(matrix, table, x+1, y)
	}

	if y != 0 && matrix[y-1][x] >= curH {
		buildTable(matrix, table, x, y-1)
	}

	if y != ly-1 && matrix[y+1][x] >= curH {
		buildTable(matrix, table, x, y+1)
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	ly := len(matrix)
	if ly == 0 {
		return [][]int{}
	}
	lx := len(matrix[0])
	if lx == 0 {
		return [][]int{}
	}

	pTable := make([][]bool, ly)
	for i := 0; i < ly; i++ {
		pTable[i] = make([]bool, lx)
	}
	for i := 0; i < lx; i++ {
		buildTable(matrix, pTable, i, 0)
	}
	for i := 0; i < ly; i++ {
		buildTable(matrix, pTable, 0, i)
	}

	aTable := make([][]bool, ly)
	for i := 0; i < ly; i++ {
		aTable[i] = make([]bool, lx)
	}
	for i := 0; i < lx; i++ {
		buildTable(matrix, aTable, i, ly-1)
	}
	for i := 0; i < ly; i++ {
		buildTable(matrix, aTable, lx-1, i)
	}

	res := [][]int{}
	for j := 0; j < ly; j++ {
		for i := 0; i < lx; i++ {
			if pTable[j][i] && aTable[j][i] {
				res = append(res, []int{j, i})
			}
		}
	}

	return res
}

// @lc code=end
