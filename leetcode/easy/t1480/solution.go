package t1480

// 1480. 一维数组的动态和
// https://leetcode.cn/problems/running-sum-of-1d-array
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
