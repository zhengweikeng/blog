package main

import (
	"fmt"
)

func insert(data []int, value int) []int {
	data = append(data, value)
	swim(data)

	return data
}

func remove(data []int, value int) []int {

	return data
}

func swim(data []int) {
	i := len(data) - 1
	for i > 1 && data[i/2] < data[i] {
		data[i/2], data[i] = data[i], data[i/2]
		i = i / 2
	}
}

func sink(data []int, n, index int) {
	for {
		maxIndex := index
		if index*2 <= n && data[index] < data[index*2] {
			maxIndex = index * 2
		}
		if index*2+1 <= n && data[maxIndex] < data[index*2+1] {
			maxIndex = index*2 + 1
		}

		if maxIndex == index {
			break
		}

		data[maxIndex], data[index] = data[index], data[maxIndex]
		index = maxIndex
	}
}

func buildHeap(data []int) {
	length := len(data)
	for i := length / 2; i >= 1; i-- {
		sink(data, length, i)
	}
}

func sort(data []int) {
	// 堆化
	buildHeap(data)

	length := len(data)
	k := length
	for k > 1 {
		data[1], data[k] = data[k], data[1]
		k--
		sink(data, k, 1)
	}
}

func printHeap(data []int) {
	fmt.Println(data[1:])
}

func main() {
	data := []int{0, 33, 17, 21, 16, 13, 15, 9, 5, 6, 7, 8, 1, 2}

	data = insert(data, 22)
	printHeap(data)
}
