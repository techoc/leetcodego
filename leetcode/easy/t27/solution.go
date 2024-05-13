package t27

// 27. 移除元素
// https://leetcode.cn/problems/remove-element
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val { // 快指针不等于移除元素时 交换快慢指针的值 慢指针前进·
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
