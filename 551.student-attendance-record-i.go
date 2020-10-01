package main

/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

// @lc code=start
func checkRecord(s string) bool {
	absent := false
	late := 0

	for _, r := range s {
		if r == 'A' {
			if absent {
				return false
			}
			absent = true
			late = 0
		} else if r == 'L' {
			late++
			if late == 3 {
				return false
			}
		} else if r == 'P' {
			late = 0
		}

	}

	return true
}

// @lc code=end
