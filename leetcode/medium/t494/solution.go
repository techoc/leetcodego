package t494

// 494. 目标和
// https://leetcode.cn/problems/target-sum/
func findTargetSumWays(nums []int, target int) (count int) {
	// 定义回溯函数，用于递归地计算所有可能的和
	var backtrack func(int, int)
	// 回溯函数，递归核心部分，用于计算给定索引位置的元素加减后与目标值的匹配情况。
	// 参数:
	//   index - 当前处理的元素索引。
	//   sum - 当前路径上的和。
	backtrack = func(index, sum int) {
		// 如果索引达到数组长度，说明遍历完所有元素
		if index == len(nums) {
			// 如果当前和等于目标值，增加计数
			if sum == target {
				count++
			}
			return
		}
		// 递归调用，分别尝试加法和减法
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	// 从数组的第一个元素开始回溯，初始和为0
	backtrack(0, 0)
	// 返回达到目标和的方法总数
	return
}
