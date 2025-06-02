package main

// 二叉树节点结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构造方法
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil, // 左节点指针
		Right: nil, // 右节点指针
		Val:   v,   // 节点值
	}
}
