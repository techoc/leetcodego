package t1470

// 1470.重新排列数组
// https://leetcode.cn/problems/shuffle-the-array/
func shuffle(nums []int, n int) []int {
	// 初始化结果数组
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		// 将两个数组的元素交替放入结果数组中
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
