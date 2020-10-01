package main

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStringsReverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func addStrings(num1 string, num2 string) string {
	a := []byte{}
	b := []byte{}
	if len(num1) > len(num2) {
		a = []byte(num1)
		b = []byte(num2)
	} else {
		a = []byte(num2)
		b = []byte(num1)
	}

	addStringsReverse(a)
	addStringsReverse(b)
	ret := []byte{}
	plus := false
	for i := 0; i < len(a); i++ {
		av := int(a[i] - '0')
		bv := 0
		if i < len(b) {
			bv = int(b[i] - '0')
		}

		v := av + bv
		if plus {
			v++
			plus = false
		}

		if v > 9 {
			v -= 10
			plus = true
		}

		ret = append(ret, byte(int('0')+v))
	}

	if plus {
		ret = append(ret, '1')
	}

	addStringsReverse(ret)

	return string(ret)
}

// @lc code=end
