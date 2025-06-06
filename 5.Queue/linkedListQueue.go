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

/* 访问队首元素 */
func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

/* 获取队列的长度 */
func (s *linkedListQueue) size() int {
	return s.data.Len()
}

/* 判断队列是否为空 */
func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListQueue) toList() *list.List {
	return s.data
}
