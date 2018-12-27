// https://leetcode-cn.com/problems/contains-duplicate/
// 对应剑指offer-第2章-面试题3：数组中重复的数字

package main

import "fmt"

/**
剑指offer的描述：
	数组中的数字都在0~n-1的范围内。如果这个数组中没有重复的数字，那么当数组排序之后数字i将出现在下标为i的位置。
	由于数组中有重复的数字，有些位置可能存在多个数字，同时有些位置可能没有数字。
*/
func containsDuplicate(nums []int) bool {
	count := len(nums)
	if count == 0 {
		return false
	}

	for i := 0; i < count; i++ {
		if nums[i] < 0 || nums[i] > count-1 {
			return false
		}
	}

	// 从头到尾扫描数组中的数字，当扫描到下标为i的数字时，首先比较这个数字（用val表示）是不是等于i。
	for i := 0; i < count; {
		j := nums[i]

		// 如果是，则接着扫描
		if j == i {
			i++
			continue
		}

		// 如果不是，比较j和第j个数字
		// 如果相等，说明重复了
		if j == nums[j] {
			return true
		}

		// 如果不相等，交换第i个数字和第j个数字
		nums[i], nums[j] = nums[j], nums[i]
	}

	return false
}

func containsDuplicateWithHash(nums []int) bool {
	hash := make(map[int]int)

	for _, num := range nums {
		if val, ok := hash[num]; ok && val == num {
			return true
		}
		hash[num] = num
	}

	return false
}

func main() {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	result := containsDuplicate(arr)
	fmt.Println(arr, "containsDuplicate:", result)

	arr = []int{3, 3}
	result = containsDuplicateWithHash(arr)
	fmt.Println(arr, "containsDuplicateWithHash:", result)
}
