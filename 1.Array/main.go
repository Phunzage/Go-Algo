// 初始化数组

package main

import "fmt"

var arr [5]int

func main() {
	// 在 Go 中，指定长度时（[5]int）为数组，不指定长度时（[]int）为切片
	// 由于 Go 的数组被设计为在编译期确定长度，因此只能使用常量来指定长度
	// 为了方便实现扩容 extend() 方法，以下将切片（Slice）看作数组（Array）
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

}
