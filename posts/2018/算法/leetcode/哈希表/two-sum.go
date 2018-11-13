// 两数之和
// https://leetcode-cn.com/problems/two-sum/description/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) (res [2]int) {
	numMap := make(map[int]int)

	for i, num := range nums {
		t := target - num
		val, ok := numMap[num]
		if ok == true {
			fmt.Println(len(res))
			res[0] = val
			res[1] = i
		} else {
			numMap[t] = i
		}
	}

	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
