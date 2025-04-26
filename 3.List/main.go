// 可以把“列表”和“动态数组”视为等同的概念

package main

import (
	"fmt"
	"sort"
)

func main() {

	// 初始化列表
	// 无初始值
	nums1 := []int{}
	// 有初始值
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println(nums1, nums)

	// 列表本质上是数组，因此可以在 O(1) 时间内访问和更新元素，效率很高
	// 访问元素
	num := nums[1] // 访问索引1处的元素
	fmt.Println(num)

	nums[1] = 0 // 将索引1处的元素改为0

	// 清空列表
	nums = nil
	fmt.Println(nums)

	// 在尾部添加元素
	nums = append(nums, 1)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 5)
	nums = append(nums, 4)
	fmt.Println(nums)

	// 在中间插入元素
	nums = append(nums[:3], append([]int{6}, nums[3:]...)...) // 在索引3处插入数字6
	// 原理分析：内部append，先创建一个切片，内含元素6，将nums索引3后面的元素加到6后面，此时新切片为[6 5 4]再将新切片加到原nums[:3]后面，为[1 3 2 6 5 4]
	// nums[3:]... 等价于将切片中的每个元素作为独立参数传递，append第二个元素需要传int元素，但是nums[3:]是一个切片，所以用...将切片解包为独立元素
	fmt.Println(nums)

	// 删除元素
	nums = append(nums[:3], nums[4:]...)
	fmt.Println(nums)

	// 通过索引遍历列表
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	// 直接遍历列表元素
	count = 0
	for _, num := range nums {
		count += num
	}

	/* 拼接两个列表 */
	nums2 := []int{6, 8, 7, 10, 9}
	nums = append(nums, nums2...) // 将列表 nums1 拼接到 nums 之后
	fmt.Println(nums)

	// 排序列表
	sort.Ints(nums)
	fmt.Println(nums)
}
