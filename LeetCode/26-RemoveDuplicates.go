package leetcode

func removeDuplicates(nums []int) int {
	a := []int{}
	for i := range nums {
		for j := 0; j < len(a); j++ {
			if i != j {
				a[j] = i
				break
			}
		}
	}
	return len(a)
}
