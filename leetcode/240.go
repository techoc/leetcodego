package leetcode

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//示例:
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
//输出：true

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for _, v := range matrix {
		//如果当前行的第一个元素大于target，则说明target不在当前行，继续查找下一行
		if v[0] > target {
			continue
		}
		//如果当前行的最后一个元素小于target，则说明target不在当前行，继续查找下一行
		if v[len(v)-1] < target {
			continue
		}
		//如果当前行的最后一个元素大于等于target，则说明target有可能在当前行，进行搜索
		if v[len(v)-1] >= target {
			for i := 0; i < len(v); i++ {
				if v[i] == target {
					return true
				}
			}
		}
	}
	//全部搜索完毕，说明target不在矩阵中
	return false
}
