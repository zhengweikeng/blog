// 链表的中间结点
// https://leetcode-cn.com/problems/middle-of-the-linked-list/

package main

import "fmt"

type linkList struct {
	value int
	next  *linkList
}

func (link *linkList) insert(value int) {
	link.next = &linkList{
		value: value,
	}
}

func (link *linkList) getMiddle() *linkList {
	slow := link
	fast := link

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func (link *linkList) print() {
	node := link
	for {
		if node == nil {
			break
		}

		fmt.Println(node.value)
		node = node.next
	}
}

func main() {
	l1 := linkList{
		value: 3,
	}
	l2 := linkList{
		value: 4,
	}
	l3 := linkList{
		value: 6,
	}
	l4 := linkList{
		value: 8,
	}
	l5 := linkList{
		value: 9,
	}
	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	// l5.next = &l3

	// fmt.Println(l1.value, l1.next.value, l1.next.next.value)
	// fmt.Println(l1.getMiddle().value)

	// revLink := reverseList(&l1)
	// revLink.print()
	// revLink := reverseListNotInPlace(&l1)
	// revLink.print()

	// fmt.Println(l1.isPalindrome())

	fmt.Println(l1.hasCycle())

	ll1 := linkList{
		value: 1,
	}
	ll2 := linkList{
		value: 4,
	}
	ll3 := linkList{
		value: 7,
	}
	ll1.next = &ll2
	ll2.next = &ll3
	mergeLink := mergeTwoLists(&l1, &ll1)
	mergeLink.print()
}
