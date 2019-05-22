package main

import "fmt"

func testBubbleSort() {
	fmt.Println("\n测试冒泡排序")
	arr := []int{3, 5, 4, 1, 2, 6}
	bubbleSort(arr, 6)
}

func testInsertSort() {
	fmt.Println("\n测试插入排序")
	arr := []int{3, 5, 4, 1, 2, 6}
	insertSort(arr, 6)
}

func main() {
	// 测试冒泡排序
	testBubbleSort()

	// 测试插入排序
	testInsertSort()
}
