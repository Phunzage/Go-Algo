// 二分查找（双开区间）

package main

func binarySearchUK(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 != r {
		m := l + (r-l)/2 //避免溢出
		if isBlue() {
			l = m
		} else {
			r = m
		}
	}
	return l / r
}

// 建模蓝红区后返回isBlue()
// e.g.
func isBlue() bool {
	return true // or false
}
