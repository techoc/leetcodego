package easy

import "sort"

func maxProduct(nums []int) int {
	//排序 找出数组里最大和第二大的数
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
