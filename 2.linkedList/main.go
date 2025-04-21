package main

// 初始化链表 1 -> 3 -> 2 -> 5 -> 4

func main() {

	// 初始化各个节点
	// 每个节点的传值是一个int类型
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)

	// 构建节点之间的引用
	// 每个结构体的Next，存放下一个指向的结构体的地址
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
}

// 数组整体是一个变量，比如数组 nums 包含元素 nums[0] 和 nums[1] 等，而链表是由多个独立的节点对象组成的。我们通常将头节点当作链表的代称，比如以上代码中的链表可记作链表 n0 。
