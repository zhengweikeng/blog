package main

import "fmt"

func selectSort(arr []int, size int) {
	fmt.Printf(" 初始状态: %v\n", arr)

	for i := 0; i < size; i++ {
		j := i
		min := j
		for ; j < size-1; j++ {
			if arr[min] > arr[j+1] {
				min = j + 1
			}
		}
		arr[i], arr[min] = arr[min], arr[i]

		fmt.Printf("第%d次插入: %v\n", i, arr)
	}
}

func testSelectSort() {
	fmt.Println("\n=======测试选择排序=======")
	arr := []int{3, 5, 4, 1, 2, 6}
	selectSort(arr, 6)
}
