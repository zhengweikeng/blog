// 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度: O(n)
// 空间复杂度: O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 存储进位
	carried := 0
	// 头结点
	result := &ListNode{
		Val: -1,
	}
	currNode := result
	p1 := l1
	p2 := l2

	for p1 != nil || p2 != nil {
		val1 := 0
		if p1 != nil {
			val1 = p1.Val
			p1 = p1.Next
		}
		val2 := 0
		if p2 != nil {
			val2 = p2.Val
			p2 = p2.Next
		}

		value := (val1 + val2 + carried) % 10
		carried = (val1 + val2 + carried) / 10
		node := &ListNode{
			Val: value,
		}
		currNode.Next = node

		currNode = currNode.Next
	}

	if carried > 0 {
		currNode.Next = &ListNode{
			Val: 1,
		}
	}

	return result.Next
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			// Next: &ListNode{
			// 	Val: 9,
			// },
		},
	}
	l2 := &ListNode{
		Val: 0,
		// Next: &ListNode{
		// 	Val: 6,
		// 	Next: &ListNode{
		// 		Val: 4,
		// 	},
		// },
	}

	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
