package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var res strings.Builder

	if (numerator < 0) != (denominator < 0) {
		res.WriteByte('-')
	}

	if numerator < 0 {
		numerator *= -1
	}

	if denominator < 0 {
		denominator *= -1
	}

	res.WriteString(strconv.Itoa(numerator / denominator))
	numerator %= denominator

	if numerator == 0 {
		return res.String()
	}

	res.WriteByte('.')

	hm := map[int]int{}
	hm[numerator] = res.Len()
	for numerator != 0 {
		numerator *= 10
		res.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator
		if index, ok := hm[numerator]; ok {
			s := res.String()
			return s[:index] + "(" + s[index:] + ")"
		}
		hm[numerator] = res.Len()
	}

	return res.String()
}

// @lc code=end
