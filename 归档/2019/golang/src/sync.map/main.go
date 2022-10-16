package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("a", 1)
	value, ok := m.Load("a")
	fmt.Println(ok, value)

	value, ok = m.LoadOrStore("a", 2)
	fmt.Println(ok, value)

	value, ok = m.LoadOrStore("b", 2)
	fmt.Println(ok, value)

	m.Delete("a")
	value, ok = m.Load("a")
	fmt.Println(ok, value)

	m.Store("c", 3)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
