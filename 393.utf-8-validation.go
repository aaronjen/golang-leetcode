package main

/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */

// @lc code=start
func validUtf8(data []int) bool {
	cnt := 0
	for _, d := range data {
		if cnt == 0 {
			if d>>5 == 0b110 {
				cnt = 1
			} else if d>>4 == 0b1110 {
				cnt = 2
			} else if d>>3 == 0b11110 {
				cnt = 3
			} else if d>>7 != 0 {
				return false
			}
		} else {
			if d>>6 == 0b10 {
				cnt--
			} else {
				return false
			}
		}

	}

	return cnt == 0
}

// @lc code=end
