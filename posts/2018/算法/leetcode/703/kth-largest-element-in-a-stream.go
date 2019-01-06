package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Val int
}
type PriorityQueue []*Item
type KthLargest struct {
	Values []int
	K      int
	Pq     PriorityQueue
}

// 构建优先队列，用最小堆，堆顶为堆中的最小元素。
func Constructor(k int, nums []int) KthLargest {
	pqSize := k
	if k > len(nums) {
		pqSize = len(nums)
	}
	pq := make(PriorityQueue, pqSize)

	for index, value := range nums[0:pqSize] {
		pq[index] = &Item{
			Val: value,
		}
	}
	heap.Init(&pq)
	if k < len(nums) && k > 0 {
		for _, value := range nums[pqSize:] {
			if value >= pq[0].Val {
				heap.Pop(&pq)
				heap.Push(&pq, value)
			}
		}
	}

	return KthLargest{
		Values: nums,
		K:      k,
		Pq:     pq,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.Values) < this.K {
		heap.Push(&this.Pq, val)
	} else if val > this.Pq[0].Val {
		heap.Pop(&this.Pq)
		heap.Push(&this.Pq, val)
	}
	this.Values = append(this.Values, val)
	return this.Pq[0].Val
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val <= pq[j].Val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, &Item{Val: item})
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	arr := []int{0}
	kthLargest := Constructor(2, arr)
	result := kthLargest.Add(-1)
	fmt.Println(result)
	result = kthLargest.Add(1)
	fmt.Println(result)
	result = kthLargest.Add(-2)
	fmt.Println(result)
	result = kthLargest.Add(-4)
	fmt.Println(result)
	result = kthLargest.Add(3)
	fmt.Println(result)
}
