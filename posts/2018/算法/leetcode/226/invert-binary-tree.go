// 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var stacks = []*TreeNode{root}
	currNode, stacks := stacks[0], stacks[1:]

	for currNode != nil {
		left := currNode.Left
		right := currNode.Right
		currNode.Left = right
		currNode.Right = left

		if left != nil {
			stacks = append(stacks, left)
		}
		if right != nil {
			stacks = append(stacks, right)
		}

		if len(stacks) > 0 {
			currNode, stacks = stacks[0], stacks[1:]
		} else {
			currNode = nil
		}
	}

	return root
}

func main() {
	tree := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right:&TreeNode{
			Val:7,
			Left:&TreeNode{
				Val:6,
			},
			Right:&TreeNode{
				Val:9,
			},
		},
	}
	invertTree(&tree)
	printTree(&tree)
}

func printTree(node *TreeNode)  {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	printTree(node.Left)
	printTree(node.Right)
}
