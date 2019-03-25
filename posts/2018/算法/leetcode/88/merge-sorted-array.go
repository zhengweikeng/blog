// 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array

// 方法一：
// 从前往后构造数组，拿nums2中的最前面的元素跟nums1中的最前面的元素比较，找到正确的排序以后插入，然后把nums1后面的元素都向后移一位。时间复杂度太高。

// 方法二：
// 新构造一个空数组nums3，那nums2中的最前面的元素跟nums1中的最前面的元素比较，然后将小的数依次插入到nums3后面。这个方法降低了时间复杂度，但是额外构造了一个数组。

// 方法三：
// 提示中已经给出，假设nums1有足够的空间了，于是我们不需要额外构造一个数组，并且可以从后面不断地比较元素进行合并。
// 1. 比较nums2与nums1中最后面的那个元素，把最大的插入第m+n位
// 2. 改变数组的索引，再次进行上面的比较，把最大的元素插入到nums1中的第m+n-1位。
// 3. 循环一直到结束。循环结束条件：当index1或index2有一个小于0时，此时就可以结束循环了。如果index2小于0，说明目的达到了。如果index1小于0，就把nums2中剩下的前面的元素都复制到nums1中去就行。
// 这里我们使用方法三。

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m+n {
		return
	}

	var i = m - 1
	var j = n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}

	if i > j {
		return
	}

	for k := 0; k <= j; k++ {
		nums1[k] = nums2[k]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{0, 0, 0}
	nums2 = []int{2, 5, 6}
	merge(nums1, 0, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1, 2, 3}
	nums2 = []int{0, 0, 0}
	merge(nums1, 3, nums2, 0)
	fmt.Println(nums1)
}
