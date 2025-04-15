// 数组插入元素

package main

// 传入一个数组，要插入的值和要插入的索引位置index
func insert(nums []int, num, index int) {
	// i 从数组最大索引位开始递减
	for i := len(nums) - 1; i > index; i-- {
		// 将前一个元素的值赋给后一个元素
		nums[i] = nums[i-1]
	}
	// 空出的index索引位赋插入的值
	nums[index] = num
}

// 在此文件中，由于数组的长度是固定的，因此插入一个元素必定会导致数组尾部元素“丢失”。
