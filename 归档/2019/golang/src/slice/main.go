package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[3:6]
	fmt.Println(len(s2), cap(s2)) // 3, 6
	fmt.Println(s2)               // 4,5,6

	s3 := s1[4:7]
	fmt.Println(s3) // 5,6,7
	s3[0] = 10
	fmt.Println(s2) // 4,10,6
	fmt.Println(s3) // 10,6,7

	s1[5] = 11
	fmt.Println(s2) // 4,10,11
	fmt.Println(s3) // 10,11,7
}
