// 初始化链表

// 链表节点 ListNode 除了包含值，还需额外保存一个指针。因此在相同数据量下，链表比数组占用更多的内存空间。

package main

// 链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一节点的指针
}

// NewListNode 构造函数，创建一个新的链表
// 传入一个值，返回一个带有值的链表
func NewListNode(val int) *ListNode {
	return &ListNode{
		// 此时值已被赋值到链表中，但是下一个节点的指针还未被赋值
		Val:  val,
		Next: nil,
	}
}
