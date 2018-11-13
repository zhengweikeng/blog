// 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/description/
package main

import (
	"fmt"
)

func isMapEqual(map1, map2 map[rune]int) bool {
	fmt.Println(map1, map2)
	if len(map1) != len(map2) {
		return false
	}

	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}

	return true
}

// 通过哈希表实现
// 时间复杂度 O(n)
func isAnagramByHash(s string, t string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, ch := range s {
		count := map1[ch]
		map1[ch] = count + 1
	}

	for _, ch := range t {
		count := map2[ch]
		map2[ch] = count + 1
	}

	return isMapEqual(map1, map2)
}

func main() {
	str1 := "anagram"
	str2 := "nagaram"
	isAnagram := isAnagramByHash(str1, str2)
	fmt.Println(isAnagram)
}
