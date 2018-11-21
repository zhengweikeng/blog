package main

import (
	"fmt"
)

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
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
		fmt.Println(arr)
	}
}

func main() {
	arr := []int{4, 5, 6, 1, 3, 2}
	insertSort(arr)
}
