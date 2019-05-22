package main

import (
	"fmt"
)

func insertSort(arr []int, size int) {
	fmt.Printf(" 初始状态: %v\n", arr)

	for i := 1; i < size; i++ {
		item := arr[i]
		j := i - 1

		// 查找插入的位置
		for ; j >= 0; j-- {
			if item < arr[j] {
				// 移动数据
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		// 插入数据
		arr[j+1] = item

		fmt.Printf("第%d次插入: %v\n", i, arr)
	}
}
