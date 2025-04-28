// 基于链表实现栈

package main

import "container/list"

// 基于链表实现的栈
type linkedListStack struct {
	// 使用内置包 list 实现栈
	data *list.List
}

// 初始化栈
func newlinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

// 入栈
func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

// 出栈
