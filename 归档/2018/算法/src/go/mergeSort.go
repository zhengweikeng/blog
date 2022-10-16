package main

import (
	"fmt"
)

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	fmt.Printf("merge前left=%d,mid=%d,right=%d: %v\n", left, mid, right, arr)
	var tmp []int
	i := left
	j := mid + 1

	// 比较左右区间的元素，将较小的元素插入临时数组中
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	// 判断哪个区间的元素还有剩余，将剩余元素添加到临时数组后面
	start := i
	end := mid
	if j <= right {
		start = j
		end = right
	}
	for start <= end {
		tmp = append(tmp, arr[start])
		start++
	}

	// 将临时数组的元素按顺序添加回原始数组
	for i := 0; i <= right-left; i++ {
		arr[left+i] = tmp[i]
	}

	fmt.Printf("merge后: %v\n\n", arr)
}

func testMergeSort() {
	fmt.Println("\n=======测试归并排序=======")
	arr := []int{3, 5, 4, 1, 2, 6}
	mergeSort(arr, 0, 5)
}
