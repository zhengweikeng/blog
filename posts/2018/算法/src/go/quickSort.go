package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition2(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	fmt.Printf("partition前start=%d,end=%d: %v\n", start, end, arr)
	pivot := arr[end]
	i := start
	j := end - 1

	for {
		for arr[i] < pivot {
			if i >= end {
				break
			}
			i++
		}

		for arr[j] > pivot {
			if j <= start {
				break
			}
			j--
		}

		// 两者相遇，终止循环
		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	// 此时的arr[i]有可能已经大于pivot，所以要交换一次
	arr[end], arr[i] = arr[i], arr[end]

	fmt.Printf("partition后pivot=%d: %v\n", i, arr)

	return i
}

func partition2(arr []int, start, end int) int {
	fmt.Printf("partition前start=%d,end=%d: %v\n", start, end, arr)
	pivot := arr[end]
	i := start
	j := start

	for ; j < end; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	fmt.Printf("partition后pivot=%d: %v\n", i, arr)
	return i
}

func testQuickSort() {
	fmt.Println("\n=======测试快速排序=======")
	// arr := []int{1, 2, 3, 4, 5, 6}
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	quickSort(arr, 0, len(arr)-1)
}
