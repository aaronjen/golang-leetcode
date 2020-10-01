package main

import "strings"

/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	tmp := []byte(S)

	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
	}

	ret := []byte{}
	cnt := 0
	for _, b := range tmp {
		if b != '-' {
			if cnt != 0 && cnt%K == 0 {
				ret = append(ret, '-')
			}
			ret = append(ret, b)
			cnt++
		}
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}

	return strings.ToUpper(string(ret))
}

// @lc code=end
