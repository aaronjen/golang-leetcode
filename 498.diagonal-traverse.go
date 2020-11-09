package main

/*
 * @lc app=leetcode id=498 lang=golang
 *
 * [498] Diagonal Traverse
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {

	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}

	i := 0
	j := 0
	isFromLeft := true
	res := []int{}
	for {
		res = append(res, matrix[i][j])

		if i == m-1 && j == n-1 {
			break
		}

		if isFromLeft {
			if j == n-1 {
				isFromLeft = false
				i++
			} else if i == 0 {
				isFromLeft = false
				j++
			} else {
				i--
				j++
			}
		} else {
			if i == m-1 {
				isFromLeft = true
				j++
			} else if j == 0 {
				isFromLeft = true
				i++
			} else {
				j--
				i++
			}
		}

	}

	return res
}

// @lc code=end
