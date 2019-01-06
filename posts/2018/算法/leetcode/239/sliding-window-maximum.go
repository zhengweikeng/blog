package main

import "fmt"

type Window struct {
	Values []int
	Index  int
	Size   int
}

func initValues(size int) *Window {
	values := make([]int, size)
	for i := range values {
		values[i] = -1
	}

	return &Window{
		Values: values,
		Index:  -1,
		Size:   size,
	}
}
func (win *Window) push(value int) {
	win.Values[win.Index+1] = value
	win.Index++
	if win.Index >= len(win.Values) {
		win.Index = 0
	}
}
func (win *Window) pop() {
	win.Values[win.Index] = -1
	win.Index--
	if win.Index < 0 {
		win.Index = -1
	}
}

func (win *Window) unshift() {
	newVal := win.Values[1:]
	win.Values = make([]int, win.Size)
	copy(win.Values, newVal)
	win.Index--
	if win.Index < 0 {
		win.Index = -1
	}
}

func (win *Window) length() int {
	return len(win.Values)
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}

	window := initValues(k)

	for i, x := range nums {
		if i >= k && window.Values[0] <= i-k {
			window.unshift()
		}

		for j := window.Index; j >= 0; j-- {
			if nums[window.Values[j]] <= x {
				window.pop()
			} else {
				break
			}
		}
		window.push(i)
		if i >= k-1 {
			result = append(result, nums[window.Values[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, -1}
	result := maxSlidingWindow(nums, 1)
	fmt.Println(result)
}
