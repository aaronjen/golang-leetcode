package main

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLifeLive(board [][]int, x, y int) bool {
	ly := len(board)
	lx := len(board[0])
	if x < 0 || y < 0 {
		return false
	}

	if x >= lx || y >= ly {
		return false
	}

	if board[y][x] == 1 || board[y][x] == 3 {
		return true
	}
	return false
}

func gameOfLifeSolve(board [][]int, x, y int) {
	liveNB := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if gameOfLifeLive(board, x-1+i, y-1+j) {
				liveNB++
			}
		}
	}
	if board[y][x] == 1 {
		if liveNB < 2 {
			board[y][x] = 3
		} else if liveNB > 3 {
			board[y][x] = 3
		}
	} else {
		if liveNB == 3 {
			board[y][x] = 2
		}
	}
}

func gameOfLife(board [][]int) {
	ly := len(board)

	if ly == 0 {
		return
	}

	lx := len(board[0])
	if lx == 0 {
		return
	}

	// 2 0->1
	// 3 1->0
	for x := 0; x < lx; x++ {
		for y := 0; y < ly; y++ {
			gameOfLifeSolve(board, x, y)
		}
	}

	for x := 0; x < lx; x++ {
		for y := 0; y < ly; y++ {
			if board[y][x] == 2 {
				board[y][x] = 1
			} else if board[y][x] == 3 {
				board[y][x] = 0
			}

		}
	}
}

// @lc code=end
