// 删除链表节点

// 删除链表的节点 n0 之后的首个节点

package main

// 传入n0
func removeItem(n0 *ListNode) {
	// 如果 n0 的下一个节点为空，直接退出
	if n0.Next == nil {
		return
	}

	// n0 -> P -> n1
	P := n0.Next
	n1 := P.Next
	// 让 n0 的下一个节点指向 n1，P无被指向节点自动被回收（Go自动垃圾处理）
	n0.Next = n1
}
