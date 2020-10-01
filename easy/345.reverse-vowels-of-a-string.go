package main

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	indices := []int{}

	for i, r := range s {
		if r == 'a' ||
			r == 'e' ||
			r == 'i' ||
			r == 'o' ||
			r == 'u' ||
			r == 'A' ||
			r == 'E' ||
			r == 'I' ||
			r == 'O' ||
			r == 'U' {
			indices = append(indices, i)
		}
	}

	n := len(indices)
	arr := []byte(s)

	for i := 0; i < n/2; i++ {
		a := indices[i]
		b := indices[n-i-1]

		arr[a], arr[b] = arr[b], arr[a]
	}

	return string(arr)
}

// @lc code=end
