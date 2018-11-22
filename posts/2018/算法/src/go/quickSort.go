package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
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

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[end], arr[i] = arr[i], arr[end]
	fmt.Println(arr, i)
	return i
}

func partition2(arr []int, start, end int) int {
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

	return i
}

func main() {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
