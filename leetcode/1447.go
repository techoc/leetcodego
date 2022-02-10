package leetcode

import "strconv"

//给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。
// 输入：n = 4
// 输出：["1/2","1/3","1/4","2/3","3/4"]
// 解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。

func SimplifiedFractions(n int) []string {
	res := make([]string, 0)
	//分子
	for i := 1; i < n; i++ {
		//分母
		for j := i + 1; j <= n; j++ {
			// 如果是最简分数，则添加到结果中
			if gcd(i, j) == 1 {
				res = append(res, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}
	return res
}

//求两数的最大公约数
//最简分数的分子和分母最大公约数为1
func gcd(i, j int) int {
	for i != j {
        if(i > j) {
            i= i-j
        }else if i < j {
            j = j-i
        }
    }
	if i==1 {
		return 1
	}
	return 0
}
