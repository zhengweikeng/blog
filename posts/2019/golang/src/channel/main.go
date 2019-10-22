package main

import (
	"fmt"
	"unsafe"
)

func printAddr(s []int) {
	fmt.Printf("printAddr: %p\n", s)
}
func printAddr2(s [1]int) {
	fmt.Printf("printAddr: %p\n", &s)
}

type User struct {
	Name string
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for item := range ch {
		fmt.Println(item)
	}

	fmt.Println(unsafe.Pointer(&ch))

	s := []int{1, 2, 3}
	fmt.Printf("%p\n", s)
	printAddr(s)

	str := [1]int{1}
	fmt.Printf("%p\n", &str)
	printAddr2(str)

	u1 := User{"a"}
	u2 := User{"b"}
	fmt.Println(u1 == u2)
}
