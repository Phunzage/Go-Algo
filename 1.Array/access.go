// 随机访问元素
package main

import "math/rand"

// 传入一个数组
func randomAccess(nums []int) (randomNum int) {
	// 定义一个随机数值，取自数组的长度 - 1（如果数组有五个元素，随机数值是0~4）
	randomIndex := rand.Intn(len(nums))
	// 返回数组[随机数值]
	randomNum = nums[randomIndex]
	return
}
