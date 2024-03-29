package t136

import "sort"

// https://leetcode.cn/problems/single-number/
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的
func singleNumber(nums []int) int {
	//排序后，相同的元素一定相邻，遍历数组，如果当前元素和下一个元素不相等，则返回当前元素
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		// 如果当前元素和下一个元素不相等，则返回当前元素
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	// 如果到达数组末尾，则返回当前元素
	return nums[len(nums)-1]
}
