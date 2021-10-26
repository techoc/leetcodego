package leetcode

func ContainsDuplicate(nums []int) bool {
	for index, value := range nums {
		for index1, value1 := range nums {
			if value1 == value && index != index1 {
				return true
			}
		}
	}
	return false
}
