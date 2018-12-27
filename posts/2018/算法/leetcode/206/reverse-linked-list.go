// 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

package main

// 在首个节点插入哨兵节点作为头结点，就地反转
func reverseList(link *linkList) *linkList {
	if link == nil {
		return nil
	}

	head := linkList{
		value: -1,
		next:  link,
	}

	prev := head.next
	curr := head.next.next
	for {
		if curr == nil {
			return head.next
		}

		prev.next = curr.next
		curr.next = head.next
		head.next = curr
		curr = prev.next
	}
}

func reverseListNotInPlace(link *linkList) *linkList {
	if link == nil {
		return nil
	}

	head := linkList{
		value: -1,
	}
	curr := link

	for curr != nil {
		nextNode := curr.next
		curr.next = head.next
		head.next = curr
		curr = nextNode
	}

	return head.next
}
