// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/description/

package main

type linkList struct {
	value int
	next  *linkList
}

func (link *linkList) insert(value int) {
	link.next = &linkList{
		value: value,
	}
}

func mergeTwoLists(link1, link2 *linkList) *linkList {
	newLink := &linkList{
		value: -1,
	}

	node1 := link1
	node2 := link2
	curr := newLink
	for node1 != nil || node2 != nil {
		if node1 == nil {
			curr.insert(node2.value)
			node2 = node2.next
		} else if node2 == nil {
			curr.insert(node1.value)
			node1 = node1.next
		} else if node1.value > node2.value {
			curr.insert(node2.value)
			node2 = node2.next
		} else {
			curr.insert(node1.value)
			node1 = node1.next
		}
		curr = curr.next
	}

	return newLink.next
}

func main() {

}
