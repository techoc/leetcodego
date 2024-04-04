package t3028

// 3028.边界上的蚂蚁
// https://leetcode.cn/problems/ant-on-the-boundary/
func returnToBoundaryCount(nums []int) int {
	var count int // 蚂蚁到达边界的次数
	var sum int   // 蚂蚁的累积步长
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 如果sum 为0，则蚂蚁到达边界
		if sum == 0 {
			count++
		}
	}
	return count
}
