package t26

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] { // 如果不相等，则慢指针前进一位，并将快指针指向的值赋给慢指针
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
