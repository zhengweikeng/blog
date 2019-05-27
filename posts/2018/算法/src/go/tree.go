package main

import (
	"fmt"
)

// Node 节点
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// TreePreOrder 前序遍历
func TreePreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%s ", node.Value)
	TreePreOrder(node.Left)
	TreePreOrder(node.Right)
}

// TreeInOrder 中序遍历
func TreeInOrder(node *Node) {
	if node == nil {
		return
	}

	TreeInOrder(node.Left)
	fmt.Printf("%s ", node.Value)
	TreeInOrder(node.Right)
}

// BstNode 二叉查找树节点
type BstNode struct {
	Value int
	Left  *BstNode
	Right *BstNode
}

// SearchInBst 在二叉搜索树中查找节点
func (node *BstNode) SearchInBst(data int) *BstNode {
	fmt.Printf("查询节点--> %d\n", data)

	currNode := node
	for currNode != nil {
		fmt.Printf("当前节点: %d\n", currNode.Value)
		if currNode.Value == data {
			fmt.Printf("命中节点: %d\n", currNode.Value)
			return currNode
		} else if currNode.Value > data {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}

	fmt.Printf("未命中节点: %d\n", data)
	return nil
}

// Insert 二叉查找树插入节点
func (node *BstNode) Insert(data int) {
	fmt.Printf("插入节点--> %d\n", data)

	currNode := node
	for currNode != nil {
		fmt.Printf("当前节点: %d\n", currNode.Value)
		// 忽略相等的情况
		if currNode.Value == data {
			fmt.Printf("插入节点: %d失败\n", data)
			return
		} else if currNode.Value > data {
			if currNode.Left == nil {
				currNode.Left = &BstNode{
					Value: data,
				}
				fmt.Printf("插入节点: %d成功\n", data)
				return
			}
			currNode = currNode.Left
		} else {
			if currNode.Right == nil {
				currNode.Right = &BstNode{
					Value: data,
				}
				fmt.Printf("插入节点: %d成功\n", data)
				return
			}
			currNode = currNode.Right
		}
	}
}

func (node *BstNode) delete(data int) bool {
	currNode := node
	var parentNode *BstNode

	for currNode != nil && currNode.Value != data {
		parentNode = currNode
		if currNode.Value > data {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}
	// 没有找到节点
	if currNode == nil {
		fmt.Printf("没有找到节点: %d\n", data)
		return false
	}
	fmt.Printf("查找到节点: %d, 父节点为: %d\n", currNode.Value, parentNode.Value)

	if currNode.Left != nil && currNode.Right != nil {
		// fmt.Print
		tmpNode := currNode.Right
		tmpParent := currNode
		for tmpNode.Left != nil {
			tmpParent = tmpNode
			tmpNode = tmpNode.Left
		}
		currNode.Value = tmpNode.Value
		currNode = tmpNode
		parentNode = tmpParent
	}

	var childNode *BstNode
	if currNode.Left != nil {
		childNode = currNode.Left
	} else if currNode.Right != nil {
		childNode = currNode.Right
	} else {
		childNode = nil
	}

	fmt.Println(parentNode.Left == currNode, parentNode.Right == currNode, childNode)
	if parentNode == nil {
		node = childNode
	} else if parentNode.Left == currNode {
		parentNode.Left = childNode
	} else {
		parentNode.Right = childNode
	}

	fmt.Printf("删除数据: %d成功\n", data)
	return true
}

func BstTreePreOrder(node *BstNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Value)
	BstTreePreOrder(node.Left)
	BstTreePreOrder(node.Right)
}

// TreePostOrder 后序遍历
func TreePostOrder(node *Node) {
	if node == nil {
		return
	}

	TreePostOrder(node.Left)
	TreePostOrder(node.Right)
	fmt.Printf("%s ", node.Value)
}

var tree = Node{
	Value: "A",
	Left: &Node{
		Value: "B",
		Left: &Node{
			Value: "D",
		},
		Right: &Node{
			Value: "E",
		},
	},
	Right: &Node{
		Value: "C",
		Left: &Node{
			Value: "F",
		},
		Right: &Node{
			Value: "G",
		},
	},
}

var bstTree = BstNode{
	Value: 33,
	Left: &BstNode{
		Value: 16,
		Left: &BstNode{
			Value: 13,
			Right: &BstNode{
				Value: 15,
			},
		},
		Right: &BstNode{
			Value: 18,
			Left: &BstNode{
				Value: 17,
			},
			Right: &BstNode{
				Value: 25,
				Left: &BstNode{
					Value: 19,
				},
				Right: &BstNode{
					Value: 27,
				},
			},
		},
	},
	Right: &BstNode{
		Value: 50,
		Left: &BstNode{
			Value: 34,
		},
		Right: &BstNode{
			Value: 58,
			Left: &BstNode{
				Value: 51,
			},
			Right: &BstNode{
				Value: 66,
			},
		},
	},
}

func testBst() {
	fmt.Println("\n=======二叉查找树查询=======")
	bstTree.SearchInBst(19)
	bstTree.SearchInBst(29)

	fmt.Println("\n=======二叉查找树插入数据=======")
	bstTree.Insert(55)
	// bstTree.Insert(55)

	fmt.Println("\n=======二叉查找树删除数据=======")
	bstTree.delete(18)
	BstTreePreOrder(&bstTree)
}

func testTreeOrder() {
	fmt.Println("\n=======前序遍历=======")
	TreePreOrder(&tree)

	fmt.Println("\n=======中序遍历=======")
	TreeInOrder(&tree)

	fmt.Println("\n=======后序遍历=======")
	TreePostOrder(&tree)

	fmt.Println()
}
