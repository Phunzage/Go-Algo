// 初始化二叉树

package main

func main() {
	// 初始化节点
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)

	// 构建节点之间的引用
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
}
