// 查找节点

// 在链表中查找值为 target 的首个节点

package main

// 传入链表和要查找的值target
func findNode(head *ListNode, target int) int {
	// 定义一个index用作索引目标的“下标”
	index := 0
	// 从头开始遍历链表
	for head != nil {
		// 找到后返回目标所对应的“下标”
		if head.Val == target {
			return index
		}

		head = head.Next
		index++
	}
	return -1
}
