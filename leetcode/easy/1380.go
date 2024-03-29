package easy

//给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。
//幸运数是指矩阵中满足同时下列两个条件的元素：
//在同一行的所有元素中最小
//在同一列的所有元素中最大
//示例 1：
//输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
//[[1,10,4,2],[9,3,8,7],[15,16,17,12]]
//输出：[15]
//解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。

func LuckyNumbers(matrix [][]int) (ans []int) {
	for _, row := range matrix {
	next:
		for j, x := range row {
			for _, y := range row {
				if y < x {
					continue next
				}
			}
			for _, r := range matrix {
				if r[j] > x {
					continue next
				}
			}
			ans = append(ans, x)
		}
	}
	return
}
