package t2908

import "math"

// 2908. 元素和最小的山形三元组 I
// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i
// minimumSum 计算 mountain 数组中三个元素组成的 mountain triplets 的最小和。
// 返回所有满足条件的 triplets 的最小和，如果不存在这样的 triplets，则返回 -1。
func minimumSum(nums []int) int {
	n := len(nums)
	// 计算每个位置右边的最小值。
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suffix[i] = min(suffix[i+1], nums[i])
	}

	ans := math.MaxInt // 初始化答案为最大整数。

	// 计算每个位置左边的最小值，并检查是否能形成 mountain triplets。
	preffix := nums[0]
	for i := 1; i < n-1; i++ {
		// 如果当前元素大于左边和右边的最小值，则更新答案。
		if nums[i] > preffix && nums[i] > suffix[i+1] {
			ans = min(ans, nums[i]+preffix+suffix[i+1])
		}
		preffix = min(preffix, nums[i])
	}
	if ans == math.MaxInt {
		return -1 // 如果未找到满足条件的 triplets，则返回 -1。
	}
	return ans
}
