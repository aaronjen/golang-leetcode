package main

/*
 * @lc app=leetcode id=475 lang=golang
 *
 * [475] Heaters
 */

// @lc code=start
func findRadius(houses []int, heaters []int) int {
	ret := 0

	for i := 0; i < len(houses); i++ {
		p := houses[i]
		min := p - heaters[0]
		if min < 0 {
			min *= -1
		}

		for j := 0; j < len(heaters); j++ {
			v := p - heaters[j]
			if v < 0 {
				v *= -1
			}

			if v < min {
				min = v
			}
		}

		if min > ret {
			ret = min
		}
	}

	return ret
}

// @lc code=end
