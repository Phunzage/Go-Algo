// 访问节点

// 访问链表中索引为 index 的节点
package main

// 传入链表，要访问的节点的“下标”
func access(head *ListNode, index int) *ListNode {
	// 开始遍历链表节点
	for i := 0; i < index; i++ {
		// 如果链表长度不足（中途遇到 head == nil），直接返回 nil，避免访问非法内存。
		if head == nil {
			return nil
		}
		// head节点依次向后遍历
		head = head.Next
	}
	// 到达index“下标”对应的节点，返回head节点
	return head
}
