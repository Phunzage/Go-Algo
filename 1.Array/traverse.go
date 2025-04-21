// 遍历数组

package main

func traverse(nums []int) {
	count := 0
	// 通过索引遍历数组
	for i := 0; i < len(nums); i++ {
		// 求和操作不是必须，只是做演示
		count += nums[i]
	}

	count = 0
	// 直接遍历数组元素
	for _, num := range nums {
		count += num
	}
	// 同时遍历数据索引和元素
	for i, num := range nums {
		count += nums[i]
		count += num
	}
}
