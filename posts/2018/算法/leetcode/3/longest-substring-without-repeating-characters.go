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
// 时间复杂度: O(len(s))
// 空间复杂度: O(len(charset))
func lengthOfLongestSubstring2(str string) int {
	freq := make([]rune, 256)
	// 左窗口边界
	left := 0
	// 右窗口边界
	right := -1
	strSize := len(str)
	res := 0

	for left < strSize {
		if right+1 < strSize && freq[str[right+1]] == 0 {
			right++
			freq[str[right]]++
		} else {
			freq[str[left]]--
			left++
		}
		res = int(math.Max(float64(res), float64(right-left+1)))
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("abcad"))
}
