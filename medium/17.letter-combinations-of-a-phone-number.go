package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	t := map[byte][]byte{
		'1': {},
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}

	ret := []string{""}

	for i := 0; i < len(digits); i++ {
		d := digits[i]

		p := t[d]
		tmp := []string{}
		for _, s := range ret {
			for _, b := range p {
				ar := []byte(s)
				ar = append(ar, b)

				tmp = append(tmp, string(ar))
			}
		}
		ret = tmp
	}

	return ret
}

// @lc code=end
