package main

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {
	table := map[byte]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]]++
	}

	type item struct {
		c byte
		n int
	}

	l := []item{}
	for c, n := range table {
		l = append(l, item{c, n})
	}

	sort.Slice(l, func(i, j int) bool {
		return l[j].n < l[i].n
	})

	var res strings.Builder

	for _, it := range l {
		for i := 0; i < it.n; i++ {
			res.WriteByte(it.c)
		}
	}

	return res.String()
}

// @lc code=end
