package main

import (
	"fmt"
)

func bubbleSort(arr []int, size int) {
	fmt.Printf(" 初始状态: %v\n", arr)

	for i := 0; i < size; i++ {
		flag := false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		fmt.Printf("第%d次冒泡: %v\n", i+1, arr)
		if !flag {
			break
		}
	}
}

func testBubbleSort() {
	fmt.Println("\n=======测试冒泡排序=======")
	arr := []int{3, 5, 4, 1, 2, 6}
	bubbleSort(arr, 6)
}
