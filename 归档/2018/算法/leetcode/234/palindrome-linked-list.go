// 回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/

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

func (link *linkList) getMiddle() *linkList {
	slow := link
	fast := link

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

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

func (link *linkList) isPalindrome() bool {
	result := true

	// 通过快慢指针获取链表中间节点
	middleLink := link.getMiddle()
	// 将中间节点之后的节点反转
	reversLink := reverseList(middleLink)

	node1 := link
	node2 := reversLink
	// 逐个遍历，直到遇到中间节点
	for node1 != middleLink {
		if node1.value != node2.value {
			result = false
		}
		node1 = node1.next
		node2 = node2.next
	}

	return result
}

func main() {

}
