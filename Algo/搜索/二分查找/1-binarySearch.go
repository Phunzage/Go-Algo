// 二分查找（双闭区间）

package main

import "fmt"

func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if target > nums[m] {
			i = m + 1
		} else if target < nums[m] {
			j = m - 1
		} else {
			return m
		}
	}
	// 未找到目标，返回-1
	return -1
}

func main() {
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	var t int
	fmt.Print("请输入要查找的数字：")
	fmt.Scanf("%d", &t)
	res := binarySearch(nums, t)
	fmt.Printf("索引位置：%d\n", res)
}
