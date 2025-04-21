// 扩容数组
package main

// 输入一个数组，要扩容的大小，返回一个扩容后的数组
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 遍历原数组，将原数组中的所有元素复制到新数组
	copy(res, nums) // copy(to, from)
	// 等价于：
	// for i, num := range nums {
	// 	res[i] = num
	// }
	// 返回新数组
	return res
}

// copy和for range都可以用来复制数组，但是copy性能更高，尤其是处理大容量数组时。
// 但是copy不能复制不同类型的数组，for range可以通过类型转换复制
