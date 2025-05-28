// 基于链表实现队列

package main

import "container/list"

type linkedListQueue struct {
	// 使用内置包 List 来实现队列
	data *list.List
}

// 初始化队列
func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

// 入队
func (s *linkedListQueue) push(value any) {
	s.data.PushBack(value)
}

/* 出队 */
func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

/* 判断队列是否为空 */
func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}
