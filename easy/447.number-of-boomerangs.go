package main

/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */

// @lc code=start
func boomerangsDis(a []int, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]

	return x*x + y*y
}

func numberOfBoomerangs(points [][]int) int {
	hashmap := map[int]int{}

	n := len(points)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			d := boomerangsDis(points[i], points[j])

			hashmap[d]++
		}

		for _, v := range hashmap {
			res += v * (v - 1)
		}
		hashmap = map[int]int{}
	}

	return res
}

// @lc code=end
