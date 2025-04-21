// 查找元素

package main

// 在数组中查找指定元素，若匹配则输出对应索引。
func find(nums []int, target int) (index int) {
	index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}
