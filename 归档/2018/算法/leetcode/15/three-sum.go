// 三数之和
// https://leetcode-cn.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	// 先排序
	sort.Ints(nums)

	for i, num := range nums {
		if num > 0 {
			break
		}

		j := i + 1
		k := len(nums) - 1
		target := 0 - num

		// 去除重复
		if i > 0 && num == nums[i-1] {
			continue
		}

		for j < k {
			if nums[j]+nums[k] == target {
				matchNums := []int{num, nums[j], nums[k]}
				result = append(result, matchNums)

				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
