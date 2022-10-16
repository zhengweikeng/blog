// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"
	"math"
)

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

// 滑动窗口
// 定义一个左右窗口的位置下标，在没有遇到重复的字符前，不断移动右窗口
// 当遇到重复在字符时，左窗口向左移动
// 移动过程中记录每次遇到的字符
// 时间复杂度: O(len(s))
// 空间复杂度: O(len(charset))
func lengthOfLongestSubstring2(str string) int {
	strSize := len(str)
	if strSize == 0 {
		return 0
	}

	freq := make([]rune, 256)
	// 左窗口边界
	left := 0
	// 右窗口边界
	right := -1
	res := 0

	for left < strSize {
		if right+1 < strSize && freq[str[right+1]] == 0 {
			freq[str[right+1]]++
			right++
		} else {
			freq[str[left]]--
			left++
		}
		res = int(math.Max(float64(res), float64(right-left+1)))
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}
