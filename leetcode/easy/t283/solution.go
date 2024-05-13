package t283

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 { // 遇到非0元素，交换到慢指针位置
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// 将后面的0填充
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
