// 反转字符串
// https://leetcode-cn.com/problems/reverse-string/

package main

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	str := "hello"
	s := []byte(str)
	reverseString(s)
	fmt.Println(string(s))
}
