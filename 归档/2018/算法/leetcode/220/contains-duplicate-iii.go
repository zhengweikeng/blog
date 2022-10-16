// https://leetcode-cn.com/problems/contains-duplicate-iii/

package main

import (
	"fmt"
	"math"
)

func containsNearbyAlmostDuplicateWithMap(nums []int, k int, t int) bool {
	hash := make(map[int]int)

	if t < 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if t == 0 {
			j, ok := hash[num]
			if ok && int(math.Abs(float64(i-j))) <= k {
				return true
			}
			hash[num] = i
		} else {
			for _, val := range hash {
				if int(math.Abs(float64(num-val))) <= t {
					return true
				}
			}
			hash[i] = num
			if len(hash) > k {
				delete(hash, i-k)
			}
		}
	}

	return false
}

func main() {
	arr := []int{1, 5, 9, 1, 5, 9}
	result := containsNearbyAlmostDuplicateWithMap(arr, 2, 3)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)

	arr = []int{-1, -1}
	result = containsNearbyAlmostDuplicateWithMap(arr, 1, -1)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)

	arr = []int{2, 2}
	result = containsNearbyAlmostDuplicateWithMap(arr, 3, 0)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)

	arr = []int{2, 1}
	result = containsNearbyAlmostDuplicateWithMap(arr, 1, 1)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)

	arr = []int{1, 2, 3, 1}
	result = containsNearbyAlmostDuplicateWithMap(arr, 3, 0)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)

	arr = []int{1, 0, 1, 1}
	result = containsNearbyAlmostDuplicateWithMap(arr, 1, 2)
	fmt.Println(arr, "containsNearbyAlmostDuplicateWithMap", result)
}
