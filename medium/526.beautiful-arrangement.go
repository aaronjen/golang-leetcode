package main

/*
 * @lc app=leetcode id=526 lang=golang
 *
 * [526] Beautiful Arrangement
 */

// @lc code=start
func countArrangementSolve(i int, t map[int]bool) int {
	if i == 1 {
		return 1
	}

	sum := 0
	for x, ok := range t {
		if !ok {
			continue
		}
		if x%i == 0 || i%x == 0 {
			newT := map[int]bool{}
			for key, val := range t {
				newT[key] = val
			}
			newT[x] = false

			sum += countArrangementSolve(i-1, newT)
		}
	}

	return sum
}

func countArrangement(N int) int {
	l := make([]int, N)
	for i := 1; i <= N; i++ {
		l[i-1] = i
	}

	t := map[int]bool{}
	for _, v := range l {
		t[v] = true
	}

	return countArrangementSolve(N, t)
}

// @lc code=end
