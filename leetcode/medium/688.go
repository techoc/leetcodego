package medium

//在一个n x n的国际象棋棋盘上，一个骑士从单元格 (row, column)开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
//象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
//https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/knight.png
//每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
//骑士继续移动，直到它走了 k 步或离开了棋盘。
//返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
//示例 1：
//输入: n = 3, k = 2, row = 0, column = 0
//输出: 0.0625
//解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
//在每一个位置上，也有两种移动可以让骑士留在棋盘上。
//骑士留在棋盘上的总概率是0.0625。

var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func KnightProbability(n, k, row, column int) float64 {
	//dp[i][j][l]表示从(i,j)开始，走了l步，骑士仍然在棋盘上的概率
	//初始化 dp[i][j][0] = 1,其他情况为0
	//转移 dp[i][j][l] = 0.125*(dp[i+1][j+2][l-1] + dp[i+1][j-2][l-1] + dp[i-1][j+2][l-1] + dp[i-1][j-2][l-1]+...)
	//结果 当l = k时 对dp[i][j][l]求和即可

	//初始化
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				//如果步数为0，则从(i,j)开始，骑士仍然在棋盘上的概率为1
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					//转移 dp[i][j][l] = 0.125*(dp[i+1][j+2][l-1] + dp[i+1][j-2][l-1] + dp[i-1][j+2][l-1] + dp[i-1][j-2][l-1]+...)
					for _, d := range dirs {
						x, y := i+d.i, j+d.j
						if 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] * 0.125
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
