package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=388 lang=golang
 *
 * [388] Longest Absolute File Path
 */

// @lc code=start
func isFile(s string) bool {
	if strings.Contains(s, ".") {
		return true
	}

	return false
}

func lengthLongestPath(input string) int {
	paths := strings.Split(input, "\n")

	folderStack := []string{}
	max := 0
	for _, path := range paths {
		depth := 0
		for depth < len(folderStack) {
			if strings.Contains(path, "\t") {
				depth++
				path = strings.Replace(path, "\t", "", 1)
			} else if strings.Contains(path, "    ") {
				depth++
				path = strings.Replace(path, "    ", "", 1)
			} else {
				break
			}
		}

		if isFile(path) {
			folderStack = folderStack[:depth]
			p := strings.Join(append(folderStack, path), "/")
			l := len(p)
			if l > max {
				max = l
			}
		} else if depth == len(folderStack) {
			folderStack = append(folderStack, path)
		} else {
			folderStack = folderStack[:depth]
			folderStack = append(folderStack, path)
		}
	}

	return max
}

// @lc code=end
