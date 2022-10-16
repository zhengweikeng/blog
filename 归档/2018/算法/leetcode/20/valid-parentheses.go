// 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/

package main

import "fmt"

const (
	s1 = 40  // (
	s2 = 41  // )
	s3 = 91  // [
	s4 = 93  // ]
	s5 = 123 // {
	s6 = 125 // }
)

func isValid(s string) bool {
	var stacks []rune
	var m = map[rune]rune{
		40:  41,
		91:  93,
		123: 125,
	}

	for _, ch := range s {
		if ch == s1 || ch == s3 || ch == s5 {
			stacks = append(stacks, ch)
			continue
		}

		if len(stacks) == 0 {
			return false
		}

		c := stacks[len(stacks)-1]
		stacks = stacks[0 : len(stacks)-1]
		if m[c] != ch {
			return false
		}
	}

	if len(stacks) > 0 {
		return false
	}

	return true
}

func main() {
	str := "]"
	result := isValid(str)
	fmt.Println(result)
}
