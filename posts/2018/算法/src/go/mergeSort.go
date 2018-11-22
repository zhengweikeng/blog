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
	var tmp []int
	i := left
	j := mid + 1

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

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

	for i := 0; i <= right-left; i++ {
		arr[left+i] = tmp[i]
	}

	fmt.Println(arr)
}

func main() {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	mergeSort(arr, 0, len(arr)-1)
}
