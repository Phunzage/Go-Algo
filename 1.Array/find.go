// 查找元素

package main

// 在数组中查找指定元素，若匹配则输出对应索引。
// 传入一个数组，要查找的元素
func find(nums []int, target int) (index int) {
	// 先定义一个index，作为目标元素的下标
	index = -1
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 数组第i个元素等于要查找的元素
		if nums[i] == target {
			// 将该下标赋给index，结束查找并返回
			index = i
			break
		}
	}
	return
}
