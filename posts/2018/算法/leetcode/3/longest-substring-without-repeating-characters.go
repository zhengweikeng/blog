// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func lengthOfLongestSubstring(str string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range str {
		lastI, ok := lastOccurred[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcad"))
}
