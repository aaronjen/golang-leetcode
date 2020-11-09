package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func sliceInsert(slice *[][]int, ind int, v []int) {
	s := *slice
	s = append(s, []int{})

	copy(s[ind+1:], s[ind:])
	s[ind] = v
	*slice = s
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[j][0] != people[i][0] {
			return people[j][0] < people[i][0]
		}

		return people[i][1] < people[j][1]
	})

	res := [][]int{}

	for _, p := range people {
		sliceInsert(&res, p[1], p)
	}

	return res
}

// @lc code=end
