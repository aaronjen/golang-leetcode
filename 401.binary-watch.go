package main

import (
	"fmt"
	"math/bits"
)

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(num int) []string {
	ret := []string{}

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			count := bits.OnesCount(uint((i << 6) | j))

			if count == num {
				ret = append(ret, fmt.Sprintf("%v:%02d", i, j))
			}
		}
	}

	return ret
}

// @lc code=end
