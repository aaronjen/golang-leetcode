package main

/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	l := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		l[i] = map[int]bool{}
	}

	for _, e := range edges {
		a, b := e[0], e[1]

		l[a][b] = true
		l[b][a] = true
	}
	leaves := []int{}
	for i, node := range l {
		if len(node) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		n -= len(leaves)
		tmp := []int{}
		for _, i := range leaves {
			for key := range l[i] {
				delete(l[key], i)
				if len(l[key]) == 1 {
					tmp = append(tmp, key)
				}
			}
		}
		leaves = tmp
	}

	return leaves
}

// @lc code=end
