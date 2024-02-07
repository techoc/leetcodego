package t2765

// 2765.最长交替子数组
// https://leetcode.cn/problems/longest-alternating-subarray
func alternatingSubarray(nums []int) int {
	res := -1
	n := len(nums)
	for firstIndex := 0; firstIndex < n; firstIndex++ {
		for i := firstIndex + 1; i < n; i++ {
			length := i - firstIndex + 1
			// 如果连续的长度是偶数，那么一定是交替的
			if nums[i]-nums[firstIndex] == (length-1)%2 {
				res = max(res, length)
			} else {
				break
			}
		}
	}
	return res
}
