package t1351

// 1351.统计有序矩阵中的负数
// https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/
func countNegatives(grid [][]int) (res int) {
	// 初始化列索引为网格最后一列
	pos := len(grid[0]) - 1
	// n 为网格的列数
	n := len(grid[0])

	// 遍历每一行
	for _, row := range grid {
		// 从当前列开始，向前搜索直到遇到非负数
		for ; pos >= 0; pos-- {
			if row[pos] >= 0 {
				break
			}
		}
		// 统计当前行负数的个数并累加到结果中
		res += n - pos - 1
	}
	return
}
