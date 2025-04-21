// 插入链表节点

// 在链表的节点 n0 之后插入节点 P

package main

// 传入插入节点前的 n0 节点和要插入的节点 P
func insertNode(n0 *ListNode, P *ListNode) {
	// n0的下一个节点是n1
	n1 := n0.Next
	// 现在让 P 的下一个节点指向 n1
	P.Next = n1
	// 最后让 n0 的下一个节点指向 P（n0即断开指向n1）
	n0.Next = P

	// n0 -> P -> n1
}

// 插入节点时间复杂度：O(1)
