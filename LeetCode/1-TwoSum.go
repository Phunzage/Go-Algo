// 两数之和
// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

// 思路：
// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，
// 如果找到了，直接返回 2 个数字的下标即可。
// 如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。

package leetcode

// 传入数组 nums 和给定的数字 target
func twoSum(nums []int, target int) []int {
	// 定义一个存放另一半数字的map
	m := map[int]int{}
	// 循环遍历数组
	for i := range nums {
		// 找出当前所需要的能组合 target 的另一半数字 定为 another
		another := target - nums[i]
		// 查看 map 中是否存在 another 这个数字
		if _, ok := m[another]; ok {
			// 如果存在，以数组形式直接返回 m[another] 的下标和 i
			// 注意第一次数字不会找到另一半，因为此时 map 中没有元素，此时if内代码不执行
			return []int{m[another], i}
		}
		// 该数字未找到另一半数字，则该数字值和对应下标存入 map，以便下次查询
		m[nums[i]] = i
	}
	return nil
}
