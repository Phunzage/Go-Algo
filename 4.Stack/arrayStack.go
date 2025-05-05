// 基于数组实现栈

package main

// 基于数组实现的栈
type arrayStack struct {
	data []int // 数据
}

// 初始化栈
func newArrayStack() *arrayStack {
	return &arrayStack{
		// 设置栈的长度为0，容量为16
		data: make([]int, 0, 16),
	}
}

// 栈的长度
func (s *arrayStack) size() int {
	return len(s.data)
}

// 栈是否为空
func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}
