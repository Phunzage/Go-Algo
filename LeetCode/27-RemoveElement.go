// 移除元素
// 给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。

// 思路：
// 这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组除去最后面的 val 数值
// 剩余元素的个数
package leetcode

// 传入数组和 val 值，返回剩余元素个数 int 类型
func removeElement(nums []int, val int) int {
	// 如果数组为空，直接返回 0
	if len(nums) == 0 {
		return 0
	}
	// 双指针，一个指向数组头部，一个指向数组尾部
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l] != val {
			l++
		}
		for l < r && nums[r] == val {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	if nums[l] == val {
		return l
	}
	return l + 1
}
