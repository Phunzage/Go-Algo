// 删除数组

package main

// 传入一个数组，要删除的索引位置index
func remove(nums []int, index int) {
	// 从索引位开始，后面的值赋给前面，索引位自然被覆盖（删除）
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}
