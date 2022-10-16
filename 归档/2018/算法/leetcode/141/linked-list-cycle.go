// 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/description/
// 这篇博客说了快慢指针判断环的理论基础：https://www.cnblogs.com/songdechiu/p/6686520.html

package main

type linkList struct {
	value int
	next  *linkList
}

func (link *linkList) hasCycle() bool {
	hasCycle := false

	slow := link
	fast := link
	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return hasCycle
}

func main() {

}
