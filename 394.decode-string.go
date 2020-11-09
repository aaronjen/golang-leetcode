package main

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	stringStack := []string{}
	numStack := []int{}

	curString := ""
	curNumber := 0
	for _, c := range s {
		if c == '[' {
			stringStack = append(stringStack, curString)
			numStack = append(numStack, curNumber)
			curString = ""
			curNumber = 0
		} else if c == ']' {
			n := len(stringStack)
			prevString := stringStack[n-1]
			repeatNumber := numStack[n-1]
			stringStack = stringStack[:n-1]
			numStack = numStack[:n-1]
			repeatString := ""
			for i := 0; i < repeatNumber; i++ {
				repeatString += curString
			}
			curString = prevString + repeatString
		} else if c >= '0' && c <= '9' {
			curNumber = curNumber*10 + int(c-'0')
		} else {
			curString += string(c)
		}
	}

	return curString
}

// @lc code=end
