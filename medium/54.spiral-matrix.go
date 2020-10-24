package main

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top := 0
	right := 0
	bottom := 0
	left := 0
	ret := []int{}
	x := 0
	y := 0
	lx := len(matrix[0])
	ly := len(matrix)
	for {
		if x >= lx-right {
			break
		}
		for ; x < lx-right; x++ {
			ret = append(ret, matrix[y][x])
		}
		x--
		top++
		y++

		if y >= ly-bottom {
			break
		}
		for ; y < ly-bottom; y++ {
			ret = append(ret, matrix[y][x])
		}
		y--
		right++
		x--

		if x <= left-1 {
			break
		}
		for ; x > left-1; x-- {
			ret = append(ret, matrix[y][x])
		}
		x++
		bottom++
		y--

		if y <= top-1 {
			break
		}
		for ; y > top-1; y-- {
			ret = append(ret, matrix[y][x])
		}
		y++
		left++
		x++
	}

	return ret
}

// @lc code=end
