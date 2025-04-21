// 扩容数组
package main

// 输入一个数组，要扩容的大小，返回一个扩容后的数组
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 遍历原数组，将原数组中的所有元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	// 返回新数组
	return res
}
