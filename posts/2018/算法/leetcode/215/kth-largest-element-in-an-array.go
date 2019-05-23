// 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

package main

import "fmt"

/**
借助快排的思想，找到一个pivot，例如第一次是数组的最后一个元素，将数组分为3个部分：
arr[0...p-1]、arr[p]、arr[p+1...N-1]

其中arr[0...p-1]>arr[p]，arr[p+1...N-1] < arr[p]

也就是说，此时下标为p的元素，是数组中第p+1大的元素。

如果`p+1=K`，那么p对应的元素就是我们要找的元素。

如果`k > p + 1`，说明要找的元素还要小于arr[p]，此时则需要从arr[p+1...N-1]查找，使用同样的方式进行查找。

如果`k < p + 1`，说明要找的元素大于arr[p]，此时需要从arr[0...p-1]处查找，同样用以上方式查找。

由于每次分区后都会选择一半的元素再继续进行分区，因此会处理元素：
n + n/2 + n/4 + ... + 1 = 2n-1

这种方式下，无需对数组进行全部排序，只是根据切分点切分数组，使数组大致有序，这种思路的时间复杂度为O(n)。
*/
func findKthLargest(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	p := end

	for true {
		p = partition(nums, start, end)

		if k == p+1 {
			break
		}

		if k > p+1 {
			start = p + 1
		} else {
			end = p - 1
		}
	}

	return nums[p]
}

func partition(nums []int, start, end int) int {
	fmt.Printf("partition前start=%d,end=%d: %v\n", start, end, nums)

	pivot := nums[end]
	i := start
	j := start

	for ; j < end; j++ {
		if nums[j] > pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	fmt.Printf("partition后pivot=%d: %v\n", i, nums)

	return i
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	result := findKthLargest(nums, 3)
	fmt.Println(result)
}
