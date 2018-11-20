// 回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/

package main

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
