// 删除有序数组中的重复项
// 给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。

// 思路：
// 双指针，一个用来遍历的指针last，一个用来寻找相同数据的指针 finder，finder先行一步，
// 只要遇到和last指向的数据不同的元素，就把last的下一个值更改为这个值，并且last进一位。
// 否则 finder 一直向下遍历直到找到不同元素或者遍历结束
// 这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数

package leetcode

// 传入数组，返回只出现一次的元素的个数 int 类型
func removeDuplicates(nums []int) int {
	// 如果数组长度为 0 ，直接返回
	if len(nums) == 0 {
		return 0
	}
	// 定义双指针，都位于数组开头
	last, finder := 0, 0
	// last遍历
	for last < len(nums)-1 {
		// 刚开始，finder 和 last 指向同一个数据，值相等，finder 进一位
		// 如果finder的值仍等于last，finder继续进位，直到找到与last值不同的元素
		for nums[finder] == nums[last] {
			finder++
			// 退出条件1，finder遍历完成，返回元素个数
			if finder == len(nums) {
				return last + 1
			}
		}
		// finder的值不等于last，将finder的值赋给last下一位，
		// 同时移动last，以新last为标准，finder继续寻找
		nums[last+1] = nums[finder]
		last++
	}
	// 退出条件二，last遍历完成，返回元素个数
	return last + 1
}
