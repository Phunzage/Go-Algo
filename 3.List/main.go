// 可以把“列表”和“动态数组”视为等同的概念

package main

import "fmt"

func main() {

	// 初始化列表
	// 无初始值
	nums1 := []int{}
	// 有初始值
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println(nums1, nums)
}
