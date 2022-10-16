// 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var nums []int
	nums = inOrder(root, nums)

	for i, num := range nums {
		if i ==0 {
			continue
		}

		if num <= nums[i-1] {
			return false
		}
	}

	return true
}

func inOrder(node *TreeNode, nums []int) []int {
	if node == nil {
		return nums
	}

	nums = inOrder(node.Left, nums)
	nums = append(nums, node.Val)
	nums = inOrder(node.Right, nums)
	return nums
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
	result := isValidBST(&tree)
	fmt.Println(result)
}
