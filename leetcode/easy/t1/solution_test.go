package t1

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	sum := twoSum(nums, 6)
	t.Log(sum)
}

func TestTwoSumHashMap(t *testing.T) {
	nums := []int{3, 2, 4}
	sum := twoSumHashMap(nums, 6)
	t.Log(sum)
}
