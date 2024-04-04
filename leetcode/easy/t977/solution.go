package t977

// 977.有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	n := len(nums)     // nums的长度
	lastNegIndex := -1 // 记录最后一个负数的索引
	// 找到最后一个负数的索引
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n) // 初始化结果切片，预分配足够的空间
	// 使用双指针法遍历切片，i指向最后一个负数，j指向第一个非负数
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		// 当i小于0时，说明左侧负数已经处理完毕，只处理右侧的非负数
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
			// 当j等于n时，说明右侧非负数已经处理完毕，只处理左侧的负数
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
			// 根据当前i和j指向的数的平方值大小，决定加入结果的是哪个数的平方
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans // 返回结果切片
}
