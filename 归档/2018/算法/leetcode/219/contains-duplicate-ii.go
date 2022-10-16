// https://leetcode-cn.com/problems/contains-duplicate-ii/

package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicateWithMap(nums []int, k int) bool {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if j, ok := hash[num]; ok && int(math.Abs(float64(i-j))) <= k {
			return true
		}
		hash[num] = i
	}

	return false
}

func main() {
	arr := []int{99, 99}
	k := 2
	result := containsNearbyDuplicateWithMap(arr, k)
	fmt.Println(arr, "containsNearbyDuplicateWithMap", result)
}
