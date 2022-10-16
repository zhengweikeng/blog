// https://leetcode-cn.com/problems/find-the-duplicate-number/

package main

import "fmt"

/**
将数字1-n从中间的数组m分为两部分，前面一般为1~m，后面一半为m+1~n
如果1~m的数字的数目超过m，说明有重复。
否则，m+1~n中肯定有重复的数字。
将上述找到的区间，再根据中间值分为两部分，重复上述步骤。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func findDuplicate(nums []int) int {
	length := len(nums)

	start := 1
	end := length - 1

	for end >= start {
		mid := (end-start)/2 + start
		count := countRange(nums, start, mid)

		if end == start {
			if count > 1 {
				return start
			}
			break
		}

		if count > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func countRange(nums []int, start, end int) int {
	count := 0
	for _, num := range nums {
		if num >= start && num <= end {
			count++
		}
	}

	return count
}

func main() {
	arr := []int{2, 3, 5, 4, 3, 2, 6, 7}
	result := findDuplicate(arr)
	fmt.Println(arr, "findDuplicate:", result)

	fmt.Println("========")

	arr = []int{2, 2, 2, 2, 2}
	result = findDuplicate(arr)
	fmt.Println(arr, "findDuplicate:", result)
}
