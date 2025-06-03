package main

// 基于环形数组实现的队列
type arrayqQueue struct {
	nums        []int // 用于存储队列元素的数组
	front       int   // 队首指针，指向队首元素
	queSize     int   // 队列长度
	queCapacity int   // 队列容量（即最大容纳元素数量）
}

// 初始化队列
func newArrayQueue(queCapacity int) *arrayqQueue {
	return &arrayqQueue{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

// 获取队列长度
func (q *arrayqQueue) size() int {
	return q.queSize
}

// 判断队列是否为空
func (q *arrayqQueue) isEmpty() bool {
	return q.queSize == 0
}

// 入队
func (q *arrayqQueue) push(num int) {
	// 当 rear == queCapacity 表示队列已满
	if q.queSize == q.queCapacity {
		return
	}
	// 计算队尾指针，指向队尾索引 + 1
	// 通过取余操作实现 rear 越过数组尾部后回到头部
	rear := (q.front + q.queSize) % q.queCapacity
	// 将 num 添加到队尾
	q.nums[rear] = rear
	q.queSize++
}

// 出队

// 访问队首元素

// 获取 Slice 用于打印
