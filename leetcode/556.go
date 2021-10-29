package leetcode

// 给你一个由二维数组 mat 表示的 m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
// 重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
// 如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
// 实例1
// 输入：mat = [[1,2],[3,4]], r = 1, c = 4
// 输出：[[1,2,3,4]]

func MatrixReshape(mat [][]int, r int, c int) [][]int {
    n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	//映射关系(i,j)→i×n+j
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = mat[i/m][i%m]
	}
	return ans
}
