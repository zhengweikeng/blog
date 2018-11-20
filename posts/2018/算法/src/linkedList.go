package main

import (
	"fmt"
)

type linkList struct {
	value int
	next  *linkList
}

func (link *linkList) insert(value int) {
	node := &linkList{
		value: value,
	}
	node.next = link.next
	link.next = node
}

// func (link *linkList) delete(data *linkList) {

// }

// func (link *linkList) query(value int) *linkList {
// 	node := &link
// 	for {
// 		if node == nil {
// 			return nil
// 		}

// 		if node.value == value {
// 			return node
// 		}

// 		node = link.next
// 	}
// }

func main() {
	link1 := linkList{
		value: 1,
	}

	link2 := linkList{
		value: 2,
	}
	link1.next = &link2

	link1.insert(3)

	fmt.Println(link1.value, link1.next.value, link1.next.next.value)
}
