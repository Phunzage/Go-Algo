package main

import "fmt"

func main() {
	fmt.Print("数组为: [2, 7, 6, 11, 5]")
	fmt.Print("输入相加值：")
	var n int
	a := []int{2, 7, 6, 11, 5}
	fmt.Scan(&n)
	res := hash(a, n)
	fmt.Print(res)
}

func hash(nums []int, target int) []int {
	table := map[int]int{}
	for i := range nums {
		another := target - nums[i]
		if _, ok := table[another]; ok {
			return []int{table[another], i}
		}
		table[nums[i]] = i
	}
	return nil
}
